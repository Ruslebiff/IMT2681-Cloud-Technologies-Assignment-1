package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func handlerNil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Default Handler: Invalid request received.") // error to console
	http.Error(w, "Invalid request", http.StatusBadRequest)   // error to http
}

func handlerDiag(w http.ResponseWriter, r *http.Request) {
	var D = &Diag{}

	// GET-requests
	gbifget, err := http.Get(GBIFRoot)
	if err != nil {
		log.Fatalln(err)
	}

	restcountriesget, err := http.Get(RestcountriesRoot)
	if err != nil {
		log.Fatalln(err)
	}

	// close connection, prevent resource leak if get-request fails
	defer gbifget.Body.Close()
	defer restcountriesget.Body.Close()

	// assign values to struct
	D.GBIFstatus = gbifget.StatusCode
	D.RestcountriesStatus = restcountriesget.StatusCode
	D.Version = "v1"
	elapsed := time.Since(StartTime)
	D.Uptime = elapsed.Seconds()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(D)
}

func handlerCountry(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/") // parts of url
	urllimit := r.URL.Query()["limit"]      // value of ?limit= query in url
	getrequestlimit := ""                   // declaration for limit in getrequest to api

	// if url contains "?limit=..."
	if urllimit != nil {
		getrequestlimit = "&limit=" + urllimit[0]
	}

	// url too short, or no contry specified after /country/
	if len(parts) <= 4 || parts[4] == "" {
		http.Error(w, "Missing valid country in url", http.StatusBadRequest)
	}

	// url ok, and has something after /country/
	if len(parts) >= 5 && parts[4] != "" {
		// structs for storing json values from api
		var C = &Country{}
		var R = &AllresultsArray{}

		// Get country info
		respcountry, err := http.Get(RestcountriesRoot + "alpha/" + parts[4] + "/")
		if err != nil {
			log.Fatalln(err)
		}

		// Close connection, prevent resource leak
		defer respcountry.Body.Close()

		// Read body data for country
		body, err := ioutil.ReadAll(respcountry.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// Parsing json for country info
		json.Unmarshal([]byte(body), C)

		// Get species info for specified country
		respspecies, err := http.Get(GBIFRoot + "occurrence/search?country=" + parts[4] + getrequestlimit)
		if err != nil {
			log.Fatalln(err)
		}

		// Close connection, prevent resource leak
		defer respspecies.Body.Close()

		// Read body data for species
		bodyspecies, err := ioutil.ReadAll(respspecies.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// parsing json for species in country
		json.Unmarshal([]byte(bodyspecies), R)

		// add species values from R into arrays in C
		for i := range R.Resnumber {
			C.Species = append(C.Species, R.Resnumber[i].Species)
			C.Specieskey = append(C.Specieskey, R.Resnumber[i].Specieskey)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(C)

	}
}

func handlerSpecies(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	var Sp = &Species{}

	if len(parts) <= 4 || parts[4] == "" {
		http.Error(w, "Missing valid species in url", http.StatusBadRequest)
	}

	if len(parts) >= 5 && parts[4] != "" {
		resp, err := http.Get(GBIFRoot + "species/" + parts[4] + "/")
		if err != nil {
			log.Fatalln(err)
		}

		respyear, err := http.Get(GBIFRoot + "species/" + parts[4] + "/name/")
		if err != nil {
			log.Fatalln(err)
		}

		// close connection, prevent resource leak
		defer resp.Body.Close()
		defer respyear.Body.Close()

		// read result for /species/*
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// json parsing for /species/*
		json.Unmarshal([]byte(body), Sp)

		// read result for /species/*/name
		bodyyear, err := ioutil.ReadAll(respyear.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// json parsing for /species/*/name
		json.Unmarshal([]byte(bodyyear), Sp)

		w.Header().Set("Content-Type", "application/json")
		if Sp.Key != 0 {
			json.NewEncoder(w).Encode(Sp)
		} else {
			fmt.Fprintln(w, "Species not found!")
		}

	}
}

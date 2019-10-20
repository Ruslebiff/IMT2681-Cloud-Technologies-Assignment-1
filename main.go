package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// GBIFRoot is root directory of GBIF API
// species example:
// http://api.gbif.org/v1/species/5231190
// http://api.gbif.org/v1/species/5231190/name for bracketYear
const GBIFRoot = "http://api.gbif.org/v1/"

// RestcountriesRoot is root directory of restcountries API
// filter example:
// https://restcountries.eu/rest/v2/alpha/no?fields=name;capital;currencies
const RestcountriesRoot = "https://restcountries.eu/rest/v2/"

// StartTime contains the timestamp when the program started
var StartTime = time.Now()

func main() {
	port := os.Getenv("PORT") // auto assign port, needed for heroku support
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handlerNil)
	http.HandleFunc("/conservation/v1/country/", handlerCountry)
	http.HandleFunc("/conservation/v1/species/", handlerSpecies)
	http.HandleFunc("/conservation/v1/diag/", handlerDiag)

	// print to console
	fmt.Println("Program started: ", StartTime)
	fmt.Println("Listening on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

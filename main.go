package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// StartTime contains the timestamp when the program started
var StartTime = time.Now()

func main() {
	port := "8080" // define http port to use

	http.HandleFunc("/", handlerNil)
	http.HandleFunc("/conservation/v1/country/", handlerCountry)
	http.HandleFunc("/conservation/v1/species/", handlerSpecies)
	http.HandleFunc("/conservation/v1/diag/", handlerDiag)

	// print to console
	fmt.Println("Program started: ", StartTime)
	fmt.Println("Listening on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Kicks off the scraping functions for the individual endpoints
	fmt.Printf("Fetching initial data...\n")
	go FetchData()

	// Endpoints
	r.HandleFunc("/api/cases/countries", GlobalCasesHandler).
		Methods("GET")
	r.HandleFunc("/api/cases/countries/{country}", CountryHandler).
		Methods("GET")
	// r.HandleFunc("/api/cases/unitedstates", USCasesHandler)

	http.Handle("/", r)

	fmt.Printf("Server running on http://localhost:8080\n\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

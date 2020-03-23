package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// countriesHandler returns the number of global cases
func countriesHandler(w http.ResponseWriter, r *http.Request) {
	// cases := getCases()
	// fmt.Printf("%d countries found", len(dataStore.globalCases))
	json.NewEncoder(w).Encode(dataStore.globalCases)
}

// countryHandler returns the stats for a specific country
func countryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryName := strings.Replace(vars["country"], "_", " ", -1)
	countryCases := findCountry(countryName)

	// Error handling if country isn't found
	if countryCases == (countryStruct{}) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Country Not Found"))
	} else {
		json.NewEncoder(w).Encode(countryCases)
	}
	// fmt.Printf("%d countries found", len(cases))
}

// summaryHandler returns the total numbers in a single object
func summaryHandler(w http.ResponseWriter, r *http.Request) {
	summary := findCountry("Total:")
	json.NewEncoder(w).Encode(summary)
}

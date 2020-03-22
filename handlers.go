package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GlobalCasesHandler returns the number of global cases
func GlobalCasesHandler(w http.ResponseWriter, r *http.Request) {
	// cases := getCases()
	// fmt.Printf("%d countries found", len(DataStore.GlobalCases))
	json.NewEncoder(w).Encode(DataStore.GlobalCases)
}

// CountryHandler returns the stats for a specific country
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryName := strings.Replace(vars["country"], "_", " ", -1)
	countryCases := FindCountry(countryName)

	// Error handling if country isn't found
	if countryCases == (CountryStruct{}) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Country Not Found"))
	} else {
		json.NewEncoder(w).Encode(countryCases)
	}
	// fmt.Printf("%d countries found", len(cases))
}

// SummaryHandler returns the total numbers in a single object
func SummaryHandler(w http.ResponseWriter, r *http.Request) {
	summary := FindCountry("Total:")
	json.NewEncoder(w).Encode(summary)
}

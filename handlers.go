package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GlobalCasesHandler returns the number of global cases
func GlobalCasesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// cases := getCases()
	// fmt.Printf("%d countries found", len(DataStore.GlobalCases))
	json.NewEncoder(w).Encode(DataStore.GlobalCases)
}

// CountryHandler returns the stats for a specific country
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	countryName := strings.Replace(vars["country"], "_", " ", -1)
	countryCases := FilterCountry(DataStore.GlobalCases, func(c CountryStruct) bool {
		return strings.Contains(strings.ToLower(c.CountryName), countryName)
	})
	// fmt.Printf("%d countries found", len(cases))
	json.NewEncoder(w).Encode(countryCases[0])
}

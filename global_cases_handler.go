package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// GlobalCasesHandler returns the number of global cases
func GlobalCasesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// cases := getCases()
	// fmt.Printf("%d countries found", len(cases))
	json.NewEncoder(w).Encode(DataStore.GlobalCases)
}

// ChinaCasesHandler returns the stats for China
func ChinaCasesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	chinaCases := FilterCountry(DataStore.GlobalCases, func(c CountryStruct) bool {
		return strings.Contains(strings.ToLower(c.CountryName), "china")
	})
	// fmt.Printf("%d countries found", len(cases))
	json.NewEncoder(w).Encode(chinaCases[0])
}

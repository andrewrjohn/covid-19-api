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
	countryCases := FilterCountry(DataStore.GlobalCases, func(c CountryStruct) bool {
		return strings.ToLower(c.CountryName) == countryName
	})

	// Error handling if country isn't found
	if len(countryCases) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Country Not Found"))
	} else {
		json.NewEncoder(w).Encode(countryCases[0])
	}
	// fmt.Printf("%d countries found", len(cases))

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// GlobalCasesHandler returns the number of global cases
func GlobalCasesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Basic request logging
	fmt.Printf("%v\n%s %s %s\n\n", time.Now(), r.Method, r.URL, r.UserAgent())

	// cases := getCases()
	// fmt.Printf("%d countries found", len(DataStore.GlobalCases))
	json.NewEncoder(w).Encode(DataStore.GlobalCases)
}

// CountryHandler returns the stats for a specific country
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Basic request logging
	fmt.Printf("%v\n%s %s %s\n\n", time.Now(), r.Method, r.URL, r.UserAgent())

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

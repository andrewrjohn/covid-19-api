package controllers

import (
	"covid-19-api/db"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type testResponse struct {
	Msg string `json:"msg"`
}

// Countries returns the number of global cases
func Countries(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.GetCountries())
}

// Country returns the stats for a specific country
func Country(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryName := strings.Replace(vars["country"], "_", " ", -1)
	countryCases := db.GetCountryByName(countryName)

	// // Error handling if country isn't found
	if countryCases == (db.CountryStruct{}) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Country Not Found"))
	} else {
		json.NewEncoder(w).Encode(countryCases)
	}
}

// Summary returns the total numbers in a single object
func Summary(w http.ResponseWriter, r *http.Request) {
	// summary := findCountry("Total:")
	// json.NewEncoder(w).Encode(summary)
	res := testResponse{"All good"}
	json.NewEncoder(w).Encode(res)
}

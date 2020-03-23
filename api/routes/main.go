// Package routes handles all the routes
package routes

import (
	"net/http"

	"covid-19-api/api/controllers"
	"covid-19-api/api/middleware"

	"github.com/gorilla/mux"
)

// Load loads all the routes
func Load() {
	r := mux.NewRouter()

	// Kicks off the scraping functions for the individual endpoints
	// fmt.Printf("Fetching initial data...\n")
	// go fetchData()

	// Middleware
	middleware.UseWithDefaults(r)

	// Endpoints
	r.HandleFunc("/api/cases/countries", controllers.Countries).
		Methods("GET")
	r.HandleFunc("/api/cases/countries/summary", controllers.Summary).
		Methods("GET")
	r.HandleFunc("/api/cases/countries/{country}", controllers.Country).
		Methods("GET")

	http.Handle("/", r)
}

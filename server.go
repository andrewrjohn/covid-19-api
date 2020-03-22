package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Kicks off the scraping functions for the individual endpoints
	fmt.Printf("Fetching initial data...\n")
	go FetchData()

	// Middleware
	r.Use(DefaultMiddleware)

	// Endpoints
	r.HandleFunc("/api/cases/countries", GlobalCasesHandler).
		Methods("GET")
	r.HandleFunc("/api/cases/countries/summary", SummaryHandler).
		Methods("GET")
	r.HandleFunc("/api/cases/countries/{country}", CountryHandler).
		Methods("GET")
	// r.HandleFunc("/api/cases/unitedstates", USCasesHandler)

	http.Handle("/", r)

	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on http://localhost%s\n\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

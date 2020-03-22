package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Server running on http://localhost:8080\n\n")

	// Endpoints
	http.Handle("/api/cases/unitedstates", http.HandlerFunc(USCasesHandler))
	http.Handle("/api/cases/global", http.HandlerFunc(GlobalCasesHandler))
	http.Handle("/api/cases/china", http.HandlerFunc(ChinaCasesHandler))

	// Kicks off the scraping functions for the individual endpoints
	go FetchData()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

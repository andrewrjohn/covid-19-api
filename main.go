package main

import (
	"covid-19-api/api/server"
)

func main() {
	// r := mux.NewRouter()

	// // Kicks off the scraping functions for the individual endpoints
	// fmt.Printf("Fetching initial data...\n")
	// go fetchData()

	// // Middleware
	// r.Use(DefaultMiddleware)

	// // Endpoints
	// r.HandleFunc("/api/cases/countries", countriesHandler).
	// 	Methods("GET")
	// r.HandleFunc("/api/cases/countries/summary", summaryHandler).
	// 	Methods("GET")
	// r.HandleFunc("/api/cases/countries/{country}", countryHandler).
	// 	Methods("GET")

	// http.Handle("/", r)

	// var port string
	// port = os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }
	// fmt.Printf("Server running on http://localhost:%s\n\n", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	server.Start()
}

// Package server handles the launching of the http server
package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"covid-19-api/api/routes"
	"covid-19-api/db"
)

// Start function starts the http server and loads in the necessary routes and middleware
func Start() {
	routes.Load()
	db.Populate()

	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on http://localhost:%s\n\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

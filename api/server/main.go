// Package server handles the launching of the http server
package server

import (
	"fmt"
	"log"
	"os"

	"covid-19-api/api/routes"
	"covid-19-api/db"

	"github.com/gin-gonic/gin"
)

// Start function starts the http server and loads in the necessary routes and middleware
func Start() {
	r := gin.Default()
	routes.Load(r)
	db.Populate()

	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on http://localhost:%s\n\n", port)
	log.Fatal(r.Run(":8080"))
}

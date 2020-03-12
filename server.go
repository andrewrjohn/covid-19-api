package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/api/cases/unitedstates", http.HandlerFunc(USCasesHandler))
	fmt.Printf("Server running on http://localhost:8080\n\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

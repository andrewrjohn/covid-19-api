package main

import (
	"log"
	"net/http"
)

// DefaultMiddleware handles request logging and setting the content-type
func DefaultMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		log.Println(r.Method, r.URL, r.UserAgent())

		next.ServeHTTP(w, r)
	})
}

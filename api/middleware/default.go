package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DefaultMiddleware handles request logging and setting the content-type
func defaultMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println(r.Method, r.URL, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}

// UseWithDefaults loads the default middleware
func UseWithDefaults(r *mux.Router) {
	r.Use(defaultMiddleware)
}

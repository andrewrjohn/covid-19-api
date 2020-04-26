package middleware

import (
	"github.com/gin-gonic/gin"
)

// DefaultMiddleware handles request logging and setting the content-type
func DefaultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// w.Header().Set("Content-Type", "application/json")
		// log.Println(r.Method, r.URL, r.UserAgent())
		// next.ServeHTTP(w, r)
		c.Next()
	}
}

// Package routes handles all the routes
package routes

import (
	"covid-19-api/api/controllers"
	"covid-19-api/api/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads all the routes
func Load(r *gin.Engine) {

	// Middleware
	r.Use(middleware.DefaultMiddleware())

	// Endpoints
	apiRoutes := r.Group("/api")
	apiRoutes.GET("/cases/countries", controllers.Countries)
	apiRoutes.GET("/cases/summary", controllers.Summary)
	apiRoutes.GET("/cases/countries/:country", controllers.Country)
}

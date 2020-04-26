package controllers

import (
	"covid-19-api/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type testResponse struct {
	Msg string `json:"msg"`
}

// Countries returns the number of global cases
func Countries(c *gin.Context) {
	c.JSON(http.StatusOK, db.GetCountries())
}

// Country returns the stats for a specific country
func Country(c *gin.Context) {
	countryName := strings.Replace(c.Param("country"), "_", " ", -1)
	countryCases := db.GetCountryByName(countryName)

	// // Error handling if country isn't found
	if countryCases == (db.CountryStruct{}) {
		c.String(http.StatusNotFound, "Country not found")
	} else {
		c.JSON(http.StatusOK, countryCases)
	}
}

// Summary returns the total numbers in a single object
func Summary(c *gin.Context) {
	summary := db.GetCountryByName("Total:")
	c.JSON(http.StatusOK, summary)
}

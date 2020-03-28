package db

import (
	"time"
)

// CountryStruct for an individual country
type CountryStruct struct {
	CountryName     string  `json:"country_name"`
	TotalCases      int64   `json:"total_cases"`
	NewCases        int64   `json:"new_cases"`
	TotalDeaths     int64   `json:"total_deaths"`
	NewDeaths       int64   `json:"new_deaths"`
	TotalRecovered  int64   `json:"total_recovered"`
	ActiveCases     int64   `json:"active_cases"`
	CriticalCases   int64   `json:"critical_cases"`
	CasesPerMillion float64 `json:"cases_per_million"`
}

type dataStruct struct {
	globalCases []CountryStruct
}

// dataStore is acting as our "DB" so we can cache it
var dataStore dataStruct

// fetchData kicks off the data scraping process
func fetchData() {
	getGlobalCases()

	// Update our cached data
	time.Sleep(10 * time.Minute)
	fetchData()
}

// Populate the cache and start the scraping schedule
func Populate() {
	go fetchData()
}

// GetCountries gets all the countries back from the cache
func GetCountries() []CountryStruct {
	return dataStore.globalCases
}

// GetCountryByName returns a single country
func GetCountryByName(name string) CountryStruct {
	return findCountry(name)
}

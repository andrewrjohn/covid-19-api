package db

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// C for an individual country
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

func getGlobalCases() {

	resp, err := http.Get("https://www.worldometers.info/coronavirus/")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	cases := []CountryStruct{}

	doc.Find("#main_table_countries_today tbody tr").Each(func(i int, s *goquery.Selection) {
		CountryName := strings.TrimSpace(s.Find("td").Eq(0).Text())

		TotalCases := stringtoInt(s.Find("td").Eq(1).Text())
		NewCases := stringtoInt(s.Find("td").Eq(2).Text())
		TotalDeaths := stringtoInt(s.Find("td").Eq(3).Text())
		NewDeaths := stringtoInt(s.Find("td").Eq(4).Text())
		TotalRecovered := stringtoInt(s.Find("td").Eq(5).Text())
		ActiveCases := stringtoInt(s.Find("td").Eq(6).Text())
		CriticalCases := stringtoInt(s.Find("td").Eq(7).Text())
		CasesPerMillion, _ := strconv.ParseFloat(s.Find("td").Eq(8).Text(), 64)

		newCountry := CountryStruct{
			CountryName,
			TotalCases,
			NewCases,
			TotalDeaths,
			NewDeaths,
			TotalRecovered,
			ActiveCases,
			CriticalCases,
			CasesPerMillion,
		}
		// fmt.Printf("%v\n", newCountry)
		cases = append(cases, newCountry)
	})
	dataStore.globalCases = cases
	// fmt.Printf(("Updated dataStore.globalCases\n"))
}

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

func stringtoInt(s string) int64 {

	var str string
	str = strings.TrimSpace(strings.Replace(strings.Replace(s, "+", "", -1), ",", "", -1))
	if str == "" {
		str = "0"
	}
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		log.Fatal((err))
	}

	return i
}

// findCountry finds a country from the dataStore based off the country name
func findCountry(name string) CountryStruct {
	var foundCountry CountryStruct
	for _, c := range dataStore.globalCases {
		if strings.ToLower(c.CountryName) == strings.ToLower(name) {
			foundCountry = c
		}
	}
	return foundCountry
}

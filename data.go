package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	GlobalCases []CountryStruct
}

// DataStore is acting as our "DB" so we can cache it
var DataStore dataStruct

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

		TotalCases := StringToInt(s.Find("td").Eq(1).Text())
		NewCases := StringToInt(s.Find("td").Eq(2).Text())
		TotalDeaths := StringToInt(s.Find("td").Eq(3).Text())
		NewDeaths := StringToInt(s.Find("td").Eq(4).Text())
		TotalRecovered := StringToInt(s.Find("td").Eq(5).Text())
		ActiveCases := StringToInt(s.Find("td").Eq(6).Text())
		CriticalCases := StringToInt(s.Find("td").Eq(7).Text())
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
	DataStore.GlobalCases = cases
	// fmt.Printf(("Updated DataStore.GlobalCases\n"))
}

// FetchData kicks off the data scraping process
func FetchData() {
	getGlobalCases()

	// Update our cached data
	time.Sleep(10 * time.Minute)
	FetchData()
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// CountryStruct for an individual country
type CountryStruct struct {
	CountryName    string `json:"country_name"`
	TotalCases     int64  `json:"total_cases"`
	TotalDeaths    int64  `json:"total_deaths"`
	TotalRecovered int64  `json:"total_recovered"`
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

	doc.Find("#main_table_countries tbody tr").Each(func(i int, s *goquery.Selection) {
		CountryName := strings.TrimSpace(s.Find("td").Eq(0).Text())

		TotalCases := StringToInt(s.Find("td").Eq(1).Text())
		TotalDeaths := StringToInt(s.Find("td").Eq(3).Text())
		TotalRecovered := StringToInt(s.Find("td").Eq(5).Text())
		newCountry := CountryStruct{
			CountryName,
			TotalCases,
			TotalDeaths,
			TotalRecovered,
		}
		// fmt.Printf("%v", newCountry)
		cases = append(cases, newCountry)
	})
	DataStore.GlobalCases = cases
	fmt.Printf(("Fetched Global Cases\n"))
}

// FetchData kicks off the data scraping process
func FetchData() {
	getGlobalCases()

	// Update our cached data
	time.Sleep(5 * time.Second)
	FetchData()
}

package db

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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
		DeathRate := strconv.FormatFloat(((float64(TotalDeaths)/float64(TotalCases))*100), 'f', 2, 64) + "%"

		newCountry := CountryStruct{
			CountryName:     CountryName,
			TotalCases:      TotalCases,
			NewCases:        NewCases,
			TotalDeaths:     TotalDeaths,
			NewDeaths:       NewDeaths,
			TotalRecovered:  TotalRecovered,
			ActiveCases:     ActiveCases,
			CriticalCases:   CriticalCases,
			CasesPerMillion: CasesPerMillion,
			DeathRate:       DeathRate,
		}
		// fmt.Printf("%v\n", newCountry)
		cases = append(cases, newCountry)
	})
	dataStore.globalCases = cases
	// fmt.Printf(("Updated dataStore.globalCases\n"))
}

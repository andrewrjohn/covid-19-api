package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type coord struct {
	CoordsName string
	Lat        float64
	Long       float64
}

func getCoords() []coord {

	resp, err := http.Get("https://www.latlong.net/category/states-236-14.html")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	coords := []coord{}

	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}
		name := s.Find("td").Eq(0).Text()
		lat := s.Find("td").Eq(1).Text()
		latF, _ := strconv.ParseFloat(lat, 32)
		long := s.Find("td").Eq(2).Text()
		longF, _ := strconv.ParseFloat(long, 32)
		newCoord := coord{
			CoordsName: name,
			Lat:        latF,
			Long:       longF,
		}
		coords = append(coords, newCoord)
	})
	return coords
}

// USCasesHandler handles the /api/cases/unitedstates route
func USCasesHandler(w http.ResponseWriter, r *http.Request) {

	// Basic request logging
	fmt.Printf("%v\n%s %s %s\n\n", time.Now(), r.Method, r.URL, r.UserAgent())

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("https://www.cdc.gov/coronavirus/2019-ncov/map-cases-us.json")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	coords := getCoords()

	data := new(USCasesDecoded)
	json.NewDecoder(resp.Body).Decode(&data)

	for i := 0; i < len(data.Data); i++ {
		state := &data.Data[i]

		for j := 0; j < len(coords); j++ {
			if strings.Contains(coords[j].CoordsName, state.Name) {
				// fmt.Printf("%d\n", state.Cases_Reported)
				state.Lat = coords[j].Lat
				state.Long = coords[j].Long
			}
		}
	}
	json.NewEncoder(w).Encode(USCasesEncoded(*data).Data)
}

package main

import (
	"log"
	"strconv"
	"strings"
)

// stringtoInt util converts strings to int64 with whitespace trimming and empty check
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
func findCountry(name string) countryStruct {
	var foundCountry countryStruct
	for _, c := range dataStore.globalCases {
		if strings.ToLower(c.CountryName) == strings.ToLower(name) {
			foundCountry = c
		}
	}
	return foundCountry
}

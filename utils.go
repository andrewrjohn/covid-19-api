package main

import (
	"log"
	"strconv"
	"strings"
)

// StringToInt util converts strings to int64 with whitespace trimming and empty check
func StringToInt(s string) int64 {

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

// FindCountry finds a country from the DataStore based off the country name
func FindCountry(name string) CountryStruct {
	var foundCountry CountryStruct
	for _, c := range DataStore.GlobalCases {
		if strings.ToLower(c.CountryName) == strings.ToLower(name) {
			foundCountry = c
		}
	}
	return foundCountry
}

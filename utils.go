package main

import (
	"log"
	"strconv"
	"strings"
)

// StringToInt util converts strings to int64 with whitespace trimming and empty check
func StringToInt(s string) int64 {

	var str string
	str = strings.TrimSpace(strings.Replace(s, ",", "", -1))
	if str == "" {
		str = "0"
	}
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		log.Fatal((err))
	}

	return i
}

// FilterCountry filters a CountryStruct from the country list
func FilterCountry(vs []CountryStruct, f func(CountryStruct) bool) []CountryStruct {
	vsf := make([]CountryStruct, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

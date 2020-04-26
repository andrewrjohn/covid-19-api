package db

import (
	"strconv"
	"strings"
)

func stringtoInt(s string) int64 {
	var str string
	str = strings.TrimSpace(strings.Replace(strings.Replace(s, "+", "", -1), ",", "", -1))
	if str == "" {
		str = "0"
	}
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}

	return i
}

func findCountry(name string) CountryStruct {
	var foundCountry CountryStruct
	for _, c := range dataStore.globalCases {
		if strings.ToLower(c.CountryName) == strings.ToLower(name) {
			foundCountry = c
		}
	}
	return foundCountry
}

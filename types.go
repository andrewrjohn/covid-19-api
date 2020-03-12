package main

// USCasesDecoded is the default format we get back from the JSON response
type USCasesDecoded struct {
	Data []struct {
		CasesReported         int     `json:"Cases Reported"`
		Name                  string  `json:"name"`
		Range                 string  `json:"range"`
		CommunityTransmission string  `json:"Community Transmission"`
		Lat                   float64 `json:"lat"`
		Long                  float64 `json:"long"`
		URL                   string
	} `json:"data"`
}

// USCasesEncoded so we can recast with the conventional naming scheme
type USCasesEncoded struct {
	Data []struct {
		CasesReported         int     `json:"cases_reported"`
		Name                  string  `json:"name"`
		Range                 string  `json:"range"`
		CommunityTransmission string  `json:"community_transmission"`
		Lat                   float64 `json:"lat"`
		Long                  float64 `json:"long"`
		URL                   string
	} `json:"data"`
}

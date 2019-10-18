package main

// Country struct from json
type Country struct {
	Code        string `json:"alpha2Code"`
	Countryname string `json:"name"`
	Countryflag string `json:"flag"`
	//Result      []Result `json:"results"`
	Species    []string //`json:"species"`
	Specieskey []int    //`json:"speciesKey"`
}

// Result array of country species
type Result struct {
	Species    string `json:"species"`
	Specieskey int    `json:"speciesKey"`
}

// AllresultsArray is a collection of all results 0, 1, 2, 3, 4 ...
type AllresultsArray struct {
	Resnumber []Result `json:"results"`
}

// Species struct from json
type Species struct {
	Key            int    `json:"key"`
	Kingdom        string `json:"kingdom"`
	Phylum         string `json:"phylum"`
	Order          string `json:"order"`
	Familiy        string `json:"family"`
	Genus          string `json:"genus"`
	ScientificName string `json:"scientificName"`
	CanonicalName  string `json:"canonicalName"`
	Year           string `json:"bracketYear"`
}

// Diag struct
type Diag struct {
	GBIFstatus          int
	RestcountriesStatus int
	Version             string
	Uptime              float64
}

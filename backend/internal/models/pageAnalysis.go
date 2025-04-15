package models

// PageAnalysis holds the structured analysis output
type PageAnalysis struct {
	HTMLVersion  string         `json:"htmlVersion"`
	Title        string         `json:"title"`
	Headings     map[string]int `json:"headings"`
	Links        LinkCounts     `json:"links"`
	HasLoginForm bool           `json:"hasLoginForm"`
}

// LinkCounts defines internal, external, inaccessible counts
type LinkCounts struct {
	Internal     int `json:"internal"`
	External     int `json:"external"`
	Inaccessible int `json:"inaccessible"`
}
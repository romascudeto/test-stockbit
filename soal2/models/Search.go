package models

type Search struct {
	Search       []MovieSearch `json:"search"`
	TotalResults string        `json:"totalResults"`
	Response     string        `json:"response"`
}

type MovieSearch struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

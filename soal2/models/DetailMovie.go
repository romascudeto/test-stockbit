package models

type Movie struct {
	Title    string `json:"title"`
	Year     string `json:"year"`
	Rated    string `json:"rated"`
	Released string `json:"released"`
	Runtime  string `json:"runtime"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Actors   string `json:"actors"`
	Plot     string `json:"plot"`
	Language string `json:"language"`
	Country  string `json:"country"`
	Awards   string `json:"awards"`
	Poster   string `json:"poster"`
	Ratings  []struct {
		Source string `json:"source"`
		Value  string `json:"value"`
	} `json:"ratings"`
	Metascore  string `json:"metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"type"`
	Dvd        string `json:"dvd"`
	BoxOffice  string `json:"boxOffice"`
	Production string `json:"production"`
	Website    string `json:"website"`
	Response   string `json:"response"`
	Error      string `json:"error"`
}

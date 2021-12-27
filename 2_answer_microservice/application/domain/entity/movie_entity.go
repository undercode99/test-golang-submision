package entity

import "time"

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbId string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type ListMovie struct {
	Movies           []Movie `json:"Search"`
	TotalMovieResult string  `json:"totalResults"`
	Response         string  `json:"Response"`
	ErrorMessage     string  `json:"Error"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
type MovieDetail struct {
	Title        string   `json:"Title"`
	Year         string   `json:"Year"`
	Rated        string   `json:"Rated"`
	Released     string   `json:"Released"`
	Runtime      string   `json:"Runtime"`
	Genre        string   `json:"Genre"`
	Director     string   `json:"Director"`
	Writer       string   `json:"Writer"`
	Actors       string   `json:"Actors"`
	Plot         string   `json:"Plot"`
	Language     string   `json:"Language"`
	Country      string   `json:"Country"`
	Awards       string   `json:"Awards"`
	Poster       string   `json:"Poster"`
	Ratings      []Rating `json:"Ratings"`
	Metascore    string   `json:"Metascore"`
	ImdbRating   string   `json:"imdbRating"`
	ImdbVotes    string   `json:"imdbVotes"`
	ImdbID       string   `json:"imdbID"`
	Type         string   `json:"Type"`
	Dvd          string   `json:"DVD"`
	BoxOffice    string   `json:"BoxOffice"`
	Production   string   `json:"Production"`
	Website      string   `json:"Website"`
	Response     string   `json:"Response"`
	ErrorMessage string   `json:"Error"`
}

type MovieSearchLog struct {
	KeywordSearch string
	CreatedAt     time.Time
}

type MovieFilter struct {
	Search string
	Page   int
}

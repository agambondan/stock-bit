package model

// Change Respon of OMBDAPI Movie / Movie Detail json to struct in
// https://mholt.github.io/json-to-go/

type MovieSearch struct {
	Search []struct {
		Title  string `json:"Title,omitempty"`
		Year   string `json:"Year,omitempty"`
		Imdbid string `json:"imdbID,omitempty"`
		Type   string `json:"Type,omitempty"`
		Poster string `json:"Poster,omitempty"`
	} `json:"Search,omitempty"`
	Totalresults string `json:"totalResults,omitempty"`
	Response     string `json:"Response,omitempty"`
}

type Movie struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"Title,omitempty"`
	Year   string `json:"Year,omitempty"`
	Imdbid string `json:"imdbID,omitempty"`
	Type   string `json:"Type,omitempty"`
	Poster string `json:"Poster,omitempty"`
}

type MovieDetail struct {
	Title    string `json:"Title,omitempty"`
	Year     string `json:"Year,omitempty"`
	Rated    string `json:"Rated,omitempty"`
	Released string `json:"Released,omitempty"`
	Runtime  string `json:"Runtime,omitempty"`
	Genre    string `json:"Genre,omitempty"`
	Director string `json:"Director,omitempty"`
	Writer   string `json:"Writer,omitempty"`
	Actors   string `json:"Actors,omitempty"`
	Plot     string `json:"Plot,omitempty"`
	Language string `json:"Language,omitempty"`
	Country  string `json:"Country,omitempty"`
	Awards   string `json:"Awards,omitempty"`
	Poster   string `json:"Poster,omitempty"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings,omitempty"`
	Metascore    string `json:"Metascore,omitempty"`
	Imdbrating   string `json:"imdbRating,omitempty"`
	Imdbvotes    string `json:"imdbVotes,omitempty"`
	Imdbid       string `json:"imdbID,omitempty"`
	Type         string `json:"Type,omitempty"`
	Totalseasons string `json:"totalSeasons,omitempty"`
	Response     string `json:"Response,omitempty"`
}

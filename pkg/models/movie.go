package models

import "time"

type Movie struct {
	Id     int    `json:"id"`
	TmdbId int    `json:"tmdb_id"`
	ImdbId string `json:"imdb_id"`

	Title       string    `json:"title"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	AgeRating   string    `json:"age_rating"`
	Overview    string    `json:"overview"`
	PosterUrl   string    `json:"poster_url"`
	Rating      float32   `json:"rating"`
	Director    string    `json:"director"`
	MainCasts   []string  `json:"main_casts"`

	Tags []string `json:"tags"`
}

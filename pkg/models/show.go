package models

import "time"

type Show struct {
	Id     int    `json:"id"`
	TmdbId int    `json:"tmdb_id"`
	ImdbId string `json:"imdb_id"`

	Title  string   `json:"title"`
	Genres []string `json:"genres"`

	FirstAirDate time.Time `json:"first_air_date"`
	LastAirDate  time.Time `json:"last_air_date"`

	Status     SeriesStatus `json:"status"`
	NoSeasons  int          `json:"no_seasons"`
	NoEpisodes int          `json:"no_episodes"`

	AgeRating string   `json:"age_rating"`
	Overview  string   `json:"overview"`
	PosterUrl string   `json:"poster_url"`
	Rating    float32  `json:"rating"`
	MainCasts []string `json:"main_casts"`

	Keywords []string `json:"keywords"`
}

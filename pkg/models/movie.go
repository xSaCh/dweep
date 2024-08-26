package models

import "time"

type FilmType string
type SeriesStatus string

const (
	Movie  FilmType = "movie"
	Series FilmType = "series"

	ReturningSeries SeriesStatus = "Returning Series"
	Ended           SeriesStatus = "Ended"
	InProduction    SeriesStatus = "In Production"
	Canceled        SeriesStatus = "Canceled"
	Planned         SeriesStatus = "Planned"
)

type FilmSeries struct {
	Film
	Status        SeriesStatus `json:"status"`
	TotalSeasons  []int        `json:"total_seasons"`
	TotalEpisodes []int        `json:"total_episodes"`
}

type Film struct {
	FilmId int    `json:"film_id"`
	TmdbId int    `json:"tmdb_id"`
	ImdbId string `json:"imdb_id"`

	Title       string    `json:"title"`
	Type        FilmType  `json:"type"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`

	Overview  string   `json:"overview"`
	PosterUrl string   `json:"poster_url"`
	Rating    float32  `json:"rating"`
	Director  string   `json:"director"`
	MainCasts []string `json:"main_casts"`

	Keywords []string `json:"keywords"`
}

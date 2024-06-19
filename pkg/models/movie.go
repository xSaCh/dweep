package models

import "time"

type MovieType string

const (
	Movie  = "movie"
	Series = "series"
)

type BasicFilm struct {
	MovId       int64     `json:"movId"`
	ImdbId      string    `json:"imdbId"`
	Title       string    `json:"title"`
	Type        MovieType `json:"type"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"releaseDate"`

	Plot      string  `json:"plot"`
	PosterUrl string  `json:"posterUrl"`
	Rating    float32 `json:"rating"`
}

func NewBasicFilm() BasicFilm {
	return BasicFilm{}
}

package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

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

type Film struct {
	FilmId int    `json:"film_id"`
	TmdbId int    `json:"tmdb_id"`
	ImdbId string `json:"imdb_id"`

	Title       string    `json:"title"`
	Type        FilmType  `json:"type"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	AgeRating   string    `json:"age_rating"`
	Overview    string    `json:"overview"`
	PosterUrl   string    `json:"poster_url"`
	Rating      float32   `json:"rating"`
	Director    string    `json:"director"`
	MainCasts   []string  `json:"main_casts"`

	Keywords []string `json:"keywords"`
}

type FilmSeries struct {
	Film
	Status        SeriesStatus `json:"status"`
	TotalSeasons  int          `json:"total_seasons"`
	TotalEpisodes int          `json:"total_episodes"`
}

type FilmMovie struct {
	Film
}

func (f *Film) String() string {
	v1 := reflect.ValueOf(*f)

	sb := strings.Builder{}
	for i := 0; i < v1.NumField(); i++ {
		field1 := v1.Field(i).Interface()
		sb.WriteString(fmt.Sprintf("'%v': %v\n", v1.Type().Field(i).Name, field1))
	}
	return sb.String()
}

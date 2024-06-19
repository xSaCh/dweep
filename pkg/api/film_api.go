package api

import (
	"time"

	"github.com/xSaCh/dweep/pkg/models"
)

type FilmApi interface {
	SearchFilmByTitle(title string, seaechFilter *SearchFilter) []models.BasicFilm
	SearchFilmByImdbId(imdbId string) models.BasicFilm
}

type SearchFilter struct {
	Type      models.MovieType
	AfterDate time.Time
}

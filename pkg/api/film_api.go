package api

import (
	"time"

	"github.com/xSaCh/dweep/pkg/models"
)

type FilmApi interface {
	SearchFilmByTitle(title string, seaechFilter *SearchFilter) []models.Film
	SearchFilmByImdbId(imdbId string) models.Film
}

type SearchFilter struct {
	Type      models.FilmType
	AfterDate time.Time
}

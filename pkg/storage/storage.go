package storage

import (
	"github.com/xSaCh/dweep/pkg/models"
)

// type StorageWatchlist interface {
// 	AddFilmToWatchlist(userId int, watchListItem D) error
// 	RemoveFilmFromWatchlist(filmId int, userId int) error
// 	GetFilmWatchlist(filmId int, userId int) (bool, error)
// }

type Storage interface {
	WatchlistStorage
}

type WatchlistStorage interface {
	AddFilm(item models.DBFilmWatchlistItem, userId int) error
	UpdateFilm(item models.DBFilmWatchlistItem, userId int) error
	RemoveFilm(filmId int, userId int) (bool, error)
	GetAllFilms(userId int) ([]models.FilmWatchlistItem, error)
	GetAllMovies(userId int) ([]models.FilmWatchlistItemMovie, error)
}

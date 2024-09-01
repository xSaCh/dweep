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
	AddMovie(item models.ReqWatchlistItemMovie, userId int) error

	// Replace all fields of the movie with the new values which are not nil/empty
	UpdateMovie(item models.ReqWatchlistItemMovie, userId int) error
	RemoveMovie(filmId int, userId int) (bool, error)
	GetAllMovies(userId int) ([]models.WatchlistItemMovie, error)
	GetMovie(filmId int, userId int) (models.WatchlistItemMovie, error)
}

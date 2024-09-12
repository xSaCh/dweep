package storage

import (
	"time"

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
	WLAddMovie(item models.ReqWatchlistItemMovie, filmId int, userId int) error

	// Replace all fields of the movie with the new values which are not nil/empty
	WLUpdateMovie(item models.ReqWatchlistItemMovie, filmId int, userId int) error
	WLRemoveMovie(filmId int, userId int) error
	WLGetAllMovies(userId int) ([]models.WatchlistItemMovie, error)
	WLGetMovie(filmId int, userId int) (models.WatchlistItemMovie, error)

	WLWatchedMovie(filmId int, userId int, watchedDate time.Time) error
}

package storage

import (
	"errors"
	"slices"
	"time"

	"github.com/xSaCh/dweep/pkg/models"
	"github.com/xSaCh/dweep/util"
)

type MemoryStore struct {
	watchlist []models.WatchlistItem
	mW        []models.WatchlistItemMovie
	sW        []models.WatchlistItemShow
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		watchlist: []models.WatchlistItem{},
	}
}

// func (m *MemoryStore) Get() []models.WatchlistItem { return m.watchlist }

// func (m *MemoryStore) AddFilm(item models.DBFilmWatchlistItem, userId int) error {

// 	fi := models.WatchlistItem{
// 		FilmId:        item.FilmId,
// 		UserId:        userId,
// 		Type:          item.Type,
// 		MyRating:      item.MyRating,
// 		MyTags:        item.MyTags,
// 		WatchStatus:   item.WatchStatus,
// 		Note:          item.Note,
// 		RecommendedBy: item.RecommendedBy,

// 		AddedOn:   time.Now(),
// 		UpdatedOn: time.Now()}

// 	if item.Type == models.TypeMovie {
// 		fi.WatchlistItemId = len(m.mW) + 1
// 		watchedD := []time.Time{}
// 		if item.WatchStatus == models.Watched {
// 			watchedD = append(watchedD, item.WatchedDate)
// 		}
// 		m.mW = append(m.mW, models.WatchlistItemMovie{
// 			WatchlistItem: fi,
// 			WatchedDates:  watchedD,
// 		})
// 	} else {
// 		return errors.New("not implemented")
// 	}
// 	return nil
// }

// func (m *MemoryStore) UpdateFilm(item models.DBFilmWatchlistItem, userId int) error {
// 	if item.Type == models.TypeShow {
// 		return errors.New("not implemented")
// 	}
// 	for i, it := range m.mW {
// 		if it.UserId == userId && it.FilmId == item.FilmId {
// 			m.mW[i].MyRating = item.MyRating
// 			m.mW[i].MyTags = item.MyTags
// 			m.mW[i].Note = item.Note
// 			m.mW[i].RecommendedBy = item.RecommendedBy
// 			m.mW[i].UpdatedOn = time.Now()

// 			if item.WatchStatus == models.Watched && m.mW[i].WatchStatus != models.Watched {
// 				m.mW[i].WatchedDates = append(m.mW[i].WatchedDates, item.WatchedDate)
// 			}

// 			m.mW[i].WatchStatus = item.WatchStatus
// 			break
// 		}
// 	}
// 	return nil
// }

// func (m *MemoryStore) RemoveFilm(filmId int, userId int) (bool, error) {
// 	for i, item := range m.mW {
// 		if item.UserId == userId && item.FilmId == filmId {
// 			m.mW = append(m.mW[:i], m.mW[i+1:]...)
// 			return true, nil
// 		}
// 	}
// 	return false, nil
// }

// func (m *MemoryStore) GetAllMovies(userId int) ([]models.WatchlistItemMovie, error) {
// 	var movies []models.WatchlistItemMovie
// 	for _, item := range m.mW {
// 		if item.UserId == userId {
// 			movies = append(movies, item)
// 		}
// 	}
// 	return movies, nil
// }

// func (m *MemoryStore) GetAllFilms(userId int) ([]models.WatchlistItem, error) {
// 	var films []models.WatchlistItem
// 	for _, item := range m.watchlist {
// 		if item.UserId == userId {
// 			films = append(films, item)
// 		}
// 	}
// 	return films, nil
// }

func (m *MemoryStore) AddMovie(item models.ReqWatchlistItemMovie, filmId int, userId int) error {
	// Check if filmID is valid or not
	// item.Id is tmdbId/ImdbId

	if filmId == 0 {
		return errors.New(util.ErrorInvalidId)
	}

	fi := models.WatchlistItem{
		FilmId:        filmId,
		Type:          models.TypeMovie,
		MyRating:      item.MyRating,
		MyTags:        item.MyTags,
		WatchStatus:   item.WatchStatus,
		Note:          item.Note,
		RecommendedBy: item.RecommendedBy,

		AddedOn:   time.Now(),
		UpdatedOn: time.Now()}

	fi.WatchlistItemId = len(m.mW) + 1
	watchedD := []time.Time{}
	if item.WatchStatus == models.Watched {
		if len(item.WatchedDates) == 0 {
			return errors.New(util.ErrorInvalidBody)
		}
		watchedD = append(watchedD, item.WatchedDates[0])
	}
	m.mW = append(m.mW, models.WatchlistItemMovie{
		WatchlistItem: fi,
		WatchedDates:  watchedD,
	})
	return nil
}

func (m *MemoryStore) UpdateMovie(item models.ReqWatchlistItemMovie, filmId int, userId int) error {
	for i, it := range m.mW {
		if it.FilmId == filmId {
			if item.MyRating != 0 {
				m.mW[i].MyRating = item.MyRating
			}
			if item.MyTags != nil {
				m.mW[i].MyTags = item.MyTags
			}
			if item.Note != "" {
				m.mW[i].Note = item.Note
			}
			if item.RecommendedBy != nil {
				m.mW[i].RecommendedBy = item.RecommendedBy
			}

			if item.WatchStatus != "" {
				m.mW[i].WatchStatus = item.WatchStatus
			}

			if item.WatchedDates != nil || len(item.WatchedDates) != 0 {
				m.mW[i].WatchedDates = item.WatchedDates
			}

			m.mW[i].UpdatedOn = time.Now()
			return nil
		}
	}
	return errors.New(util.ErrorFilmNotFound)
}

func (m *MemoryStore) RemoveMovie(filmId int, userId int) (bool, error) {
	for i, item := range m.mW {
		if item.FilmId == filmId {
			m.mW = append(m.mW[:i], m.mW[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New(util.ErrorFilmNotFound)
}

func (m *MemoryStore) GetAllMovies(userId int) ([]models.WatchlistItemMovie, error) {
	// var movies []models.WatchlistItemMovie
	// for _, item := range m.mW {
	// 	if item.UserId == userId {
	// 		movies = append(movies, item)
	// 	}
	// }
	return slices.Clone(m.mW), nil
}

func (m *MemoryStore) GetMovie(filmId int, userId int) (models.WatchlistItemMovie, error) {
	for _, item := range m.mW {
		if item.FilmId == filmId {
			return item, nil
		}
	}
	return models.WatchlistItemMovie{}, errors.New(util.ErrorFilmNotFound)
}

package storage

import (
	"errors"
	"time"

	"github.com/xSaCh/dweep/pkg/models"
)

type MemoryStore struct {
	watchlist []models.FilmWatchlistItem
	mW        []models.FilmWatchlistItemMovie
	sW        []models.FilmWatchlistItemSeries
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		watchlist: []models.FilmWatchlistItem{},
	}
}

func (m *MemoryStore) Get() []models.FilmWatchlistItem { return m.watchlist }

func (m *MemoryStore) AddFilm(item models.DBFilmWatchlistItem, userId int) error {

	fi := models.FilmWatchlistItem{
		FilmId:        item.FilmId,
		UserId:        userId,
		Type:          item.Type,
		MyRating:      item.MyRating,
		MyTags:        item.MyTags,
		WatchStatus:   item.WatchStatus,
		Note:          item.Note,
		RecommendedBy: item.RecommendedBy,

		AddedOn:   time.Now(),
		UpdatedOn: time.Now()}

	if item.Type == models.TypeMovie {
		fi.WatchlistItemId = len(m.mW) + 1
		watchedD := []time.Time{}
		if item.WatchStatus == models.Watched {
			watchedD = append(watchedD, item.WatchedDate)
		}
		m.mW = append(m.mW, models.FilmWatchlistItemMovie{
			FilmWatchlistItem: fi,
			WatchedDates:      watchedD,
		})
	} else {
		return errors.New("not implemented")
	}
	return nil
}

func (m *MemoryStore) UpdateFilm(item models.DBFilmWatchlistItem, userId int) error {
	if item.Type == models.TypeShow {
		return errors.New("not implemented")
	}
	for i, it := range m.mW {
		if it.UserId == userId && it.FilmId == item.FilmId {
			m.mW[i].MyRating = item.MyRating
			m.mW[i].MyTags = item.MyTags
			m.mW[i].Note = item.Note
			m.mW[i].RecommendedBy = item.RecommendedBy
			m.mW[i].UpdatedOn = time.Now()

			if item.WatchStatus == models.Watched && m.mW[i].WatchStatus != models.Watched {
				m.mW[i].WatchedDates = append(m.mW[i].WatchedDates, item.WatchedDate)
			}

			m.mW[i].WatchStatus = item.WatchStatus
			break
		}
	}
	return nil
}

func (m *MemoryStore) RemoveFilm(filmId int, userId int) (bool, error) {
	for i, item := range m.mW {
		if item.UserId == userId && item.FilmId == filmId {
			m.mW = append(m.mW[:i], m.mW[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (m *MemoryStore) GetAllMovies(userId int) ([]models.FilmWatchlistItemMovie, error) {
	var movies []models.FilmWatchlistItemMovie
	for _, item := range m.mW {
		if item.UserId == userId {
			movies = append(movies, item)
		}
	}
	return movies, nil
}

func (m *MemoryStore) GetAllFilms(userId int) ([]models.FilmWatchlistItem, error) {
	var films []models.FilmWatchlistItem
	for _, item := range m.watchlist {
		if item.UserId == userId {
			films = append(films, item)
		}
	}
	return films, nil
}

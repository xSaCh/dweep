package storage_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xSaCh/dweep/pkg/mocks"
	"github.com/xSaCh/dweep/pkg/models"
	"github.com/xSaCh/dweep/pkg/storage"
)

func WatchlistToReq(w models.WatchlistItemMovie) (models.ReqWatchlistItemMovie, int) {
	return models.ReqWatchlistItemMovie{
		ReqWatchlistItem: models.ReqWatchlistItem{

			MyRating:      w.MyRating,
			MyTags:        w.MyTags,
			WatchStatus:   w.WatchStatus,
			Note:          w.Note,
			RecommendedBy: w.RecommendedBy},
		WatchedDates: w.WatchedDates}, w.FilmId
}

func testStore_AddMovie(t *testing.T, ms storage.Storage) {

	w1 := models.WatchlistItemMovie{
		WatchlistItem: models.WatchlistItem{
			FilmId:        mocks.MovieFilms[0].FilmId,
			MyRating:      4,
			MyTags:        []string{"t1", "t2"},
			WatchStatus:   models.Watched,
			Note:          "Notes to test",
			RecommendedBy: []int64{69, 420},
		},
		WatchedDates: []time.Time{time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC)},
	}
	w2 := models.WatchlistItemMovie{
		WatchlistItem: models.WatchlistItem{
			FilmId:      mocks.MovieFilms[1].FilmId,
			WatchStatus: models.PlanToWatch,
			Note:        "Notes to test",
		},
	}

	f1, f1I := WatchlistToReq(w1)
	f2, f2I := WatchlistToReq(w2)

	ms.AddMovie(f1, f1I, 1)
	ms.AddMovie(f2, f2I, 1)
	m, err := ms.GetAllMovies(1)

	assert.NoError(t, err)
	assert.Len(t, m, 2)
	assert.Equal(t, w1.WatchedDates, m[0].WatchedDates)
	assert.Empty(t, w2.WatchedDates)
}

func testStore_UpdateMovie(t *testing.T, ms storage.Storage) {

	w1 := models.WatchlistItemMovie{
		WatchlistItem: models.WatchlistItem{
			FilmId:        mocks.MovieFilms[0].FilmId,
			MyRating:      4,
			MyTags:        []string{"t1", "t2"},
			WatchStatus:   models.Watched,
			Note:          "Notes to test",
			RecommendedBy: []int64{69, 420},
		},
		WatchedDates: []time.Time{time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC)},
	}
	w2 := models.WatchlistItemMovie{
		WatchlistItem: models.WatchlistItem{
			FilmId:      mocks.MovieFilms[1].FilmId,
			WatchStatus: models.PlanToWatch,
			Note:        "Notes to test",
		},
	}

	f1, f1I := WatchlistToReq(w1)

	f2, f2I := WatchlistToReq(w2)

	ms.AddMovie(f1, f1I, 1)
	ms.AddMovie(f2, f2I, 1)

	w1.WatchedDates = append(w1.WatchedDates, time.Date(2024, 9, 9, 0, 0, 0, 0, time.UTC))
	f1.WatchedDates = append(f1.WatchedDates, time.Date(2024, 9, 9, 0, 0, 0, 0, time.UTC))

	w2.WatchStatus = models.Dropped

	ms.UpdateMovie(f1, f1I, 1)
	ms.UpdateMovie(f2, f2I, 1)

	m, err := ms.GetAllMovies(1)

	assert.Len(t, m, 2)
	assert.NoError(t, err)
	assert.Equal(t, w1.WatchedDates, m[0].WatchedDates)
	assert.Empty(t, w2.WatchedDates, m[1].WatchedDates)
}

func testStore_RemoveMovie(t *testing.T, ms storage.Storage) {

	w1 := models.WatchlistItemMovie{
		WatchlistItem: models.WatchlistItem{
			FilmId:        mocks.MovieFilms[0].FilmId,
			MyRating:      4,
			MyTags:        []string{"t1", "t2"},
			WatchStatus:   models.Watched,
			Note:          "Notes to test",
			RecommendedBy: []int64{69, 420},
		},
		WatchedDates: []time.Time{time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC)},
	}
	w2 := models.WatchlistItemMovie{
		WatchlistItem: models.WatchlistItem{
			FilmId:      mocks.MovieFilms[1].FilmId,
			WatchStatus: models.PlanToWatch,
			Note:        "Notes to test",
		},
	}

	f1, f1I := WatchlistToReq(w1)

	f2, f2I := WatchlistToReq(w2)

	ms.AddMovie(f1, f1I, 1)
	ms.AddMovie(f2, f2I, 1)

	ms.RemoveMovie(w1.FilmId, 1)

	m, err := ms.GetAllMovies(1)

	assert.Len(t, m, 1)
	assert.NoError(t, err)
	assert.Equal(t, w2.FilmId, m[0].FilmId)
}

func TestMemoryStore(t *testing.T) {
	ms := storage.NewMemoryStore()

	t.Run("MemoryStore", func(t *testing.T) {
		testStore_AddMovie(t, ms)
		ms = storage.NewMemoryStore()
		testStore_UpdateMovie(t, ms)
		ms = storage.NewMemoryStore()
		testStore_RemoveMovie(t, ms)
	})

	ss, err := storage.NewSqlliteStore("test.db")
	assert.NoError(t, err)

	assert.NoError(t, ss.Create())

	t.Run("SqliteStore", func(t *testing.T) {
		testStore_AddMovie(t, ss)
		// ms = storage.NewMemoryStore()
		// testStore_UpdateMovie(t, ms)
		// ms = storage.NewMemoryStore()
		// testStore_RemoveMovie(t, ms)
	})
	// testStore_AddMovie(t, ms)
	// ms = storage.NewMemoryStore()
	// testStore_UpdateMovie(t, ms)
	// ms = storage.NewMemoryStore()
	// testStore_RemoveMovie(t, ms)
}

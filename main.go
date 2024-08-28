package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/xSaCh/dweep/pkg/mocks"
	"github.com/xSaCh/dweep/pkg/models"
	"github.com/xSaCh/dweep/pkg/storage"
)

const (
	OMDB_API_KEY = "5d0b46f6"
	TMDB_API_KEY = "a3ca43df787ec6b692b7e1e2d53a65ec"
)

func main() {

	// ta := api.NewTmdbApi(TMDB_API_KEY)
	// m := *ta.GetSeries("65784")
	// m := *ta.GetMovie("534780")
	memStore := storage.NewMemoryStore()

	f1 := models.DBFilmWatchlistItem{
		FilmId:        mocks.MovieFilms[0].FilmId,
		Type:          mocks.MovieFilms[0].Type,
		MyRating:      4,
		MyTags:        []string{},
		WatchStatus:   models.Watched,
		Note:          "",
		RecommendedBy: []int64{},
		WatchedDate:   time.Now(),
	}
	f2 := models.DBFilmWatchlistItem{
		FilmId:        mocks.MovieFilms[2].FilmId,
		Type:          mocks.MovieFilms[2].Type,
		MyRating:      0,
		MyTags:        []string{"Q"},
		WatchStatus:   models.PlanToWatch,
		Note:          "N",
		RecommendedBy: []int64{2},
	}

	var st storage.Storage = memStore

	st.AddFilm(f1, 1)
	st.AddFilm(f2, 1)
	f, _ := st.GetAllMovies(1)
	b, _ := json.MarshalIndent(f[1], "", "  ")
	fmt.Println(string(b))

	f2.WatchStatus = models.Watched
	f2.MyRating = 5
	st.UpdateFilm(f2, 1)
	st.RemoveFilm(f1.FilmId, 1)
	f, _ = st.GetAllMovies(1)
	b, _ = json.MarshalIndent(f, "", "  ")
	fmt.Println(string(b))

}

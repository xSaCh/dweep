package main

import (
	"encoding/json"
	"fmt"
	"time"

	ext "github.com/xSaCh/dweep/external_apis"
	"github.com/xSaCh/dweep/pkg"
	"github.com/xSaCh/dweep/pkg/mocks"
	"github.com/xSaCh/dweep/pkg/models"
	"github.com/xSaCh/dweep/pkg/storage"
	"github.com/xSaCh/dweep/util"
)

const (
	OMDB_API_KEY = "5d0b46f6"
	TMDB_API_KEY = "a3ca43df787ec6b692b7e1e2d53a65ec"
)

func main() {
	ms := storage.NewMemoryStore()

	f1 := models.ReqWatchlistItemMovie{
		ReqWatchlistItem: models.ReqWatchlistItem{
			MyRating:      4,
			MyTags:        []string{},
			WatchStatus:   models.Watched,
			Note:          "",
			RecommendedBy: []int64{},
		},
		WatchedDates: []time.Time{time.Now()},
	}

	b, _ := json.MarshalIndent(f1, "", "  ")
	fmt.Println(string(b))

	ms.AddMovie(f1, mocks.MovieFilms[0].FilmId, 1)
	ser := pkg.NewAPIServer("localhost:8080", ms)
	ser.Run()
}

func amain() {

	ta := ext.NewTmdbApi(TMDB_API_KEY)
	// m := *ta.GetMovie("550")
	m := *ta.GetSeries("65784")
	a, _ := json.Marshal(m)
	fmt.Printf("%v\n", util.MissingStructFields(m))
	fmt.Printf("%s\n\n\n\n", string(a))
	// // m := *ta.GetMovie("534780")
	// memStore := storage.NewMemoryStore()

	// f1 := models.ReqWatchlistItemMovie{
	// 	ReqWatchlistItem: models.ReqWatchlistItem{
	// 		Id:            mocks.MovieFilms[0].FilmId,
	// 		MyRating:      4,
	// 		MyTags:        []string{},
	// 		WatchStatus:   models.Watched,
	// 		Note:          "",
	// 		RecommendedBy: []int64{}},
	// 	WatchedDate: time.Now(),
	// }
	// f2 := models.DBFilmWatchlistItem{
	// 	FilmId:        mocks.MovieFilms[2].FilmId,
	// 	Type:          mocks.MovieFilms[2].Type,
	// 	MyRating:      0,
	// 	MyTags:        []string{"Q"},
	// 	WatchStatus:   models.PlanToWatch,
	// 	Note:          "N",
	// 	RecommendedBy: []int64{2},
	// }

	// var st storage.Storage = memStore

	// st.AddFilm(f1, 1)
	// st.AddFilm(f2, 1)
	// f, _ := st.GetAllMovies(1)
	// b, _ := json.MarshalIndent(f[1], "", "  ")
	// fmt.Println(string(b))

	// f2.WatchStatus = models.Watched
	// f2.MyRating = 5
	// st.UpdateFilm(f2, 1)
	// st.RemoveFilm(f1.FilmId, 1)
	// f, _ = st.GetAllMovies(1)
	// b, _ = json.MarshalIndent(f, "", "  ")
	// fmt.Println(string(b))

}

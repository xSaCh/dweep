package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/xSaCh/dweep/pkg/models"
)

const (
	TMDB_BASE_URL = "https://api.themoviedb.org/3/"
	// imageUrl = "https://image.tmdb.org/t/p/"
)

type TmdbApi struct {
	apiKey string
}

func NewTmdbApi(apiKey string) *TmdbApi {
	return &TmdbApi{apiKey: apiKey}
}

func (api *TmdbApi) SearchFilmByTitle(title string, seaechFilter *SearchFilter) []models.BasicFilm {
	print("Getting... ", TMDB_BASE_URL+"/search/movie?api_key="+api.apiKey+"&query="+title)
	res, err := http.Get(TMDB_BASE_URL + "/search/movie?api_key=" + api.apiKey + "&query=" + title)
	if err != nil {
		return []models.BasicFilm{}
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	var jsData map[string]interface{}
	json.Unmarshal(data, &jsData)

	// var films []models.BasicFilm
	// for _, v := range jsData["results"].([]interface{})  {
	// 	films = append(films, models.BasicFilm{
	// 		// ImdbId: v['i'],
	// 	})
	// }
	return []models.BasicFilm{}

}

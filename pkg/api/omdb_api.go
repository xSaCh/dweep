package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/xSaCh/dweep/pkg/models"
)

const (
	OMDB_BASE_URL = "http://www.omdbapi.com/"
	// imageUrl = "https://image.tmdb.org/t/p/"
)

type OmdbApi struct {
	apiKey string
}

func NewOmdbApi(apiKey string) *OmdbApi {
	return &OmdbApi{apiKey: apiKey}
}

func (api *OmdbApi) SearchFilmByTitle(queryTitle string, seaechFilter *SearchFilter) []models.BasicFilm {
	print("Getting... ", OMDB_BASE_URL+"?apikey="+api.apiKey+"&s="+queryTitle)

	res, err := http.Get(OMDB_BASE_URL + "?apikey=" + api.apiKey + "&s=" + queryTitle)
	if err != nil {
		return []models.BasicFilm{}
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	var jsData map[string]any
	json.Unmarshal(data, &jsData)

	var films []models.BasicFilm
	for _, v := range jsData["Search"].([]any) {
		data := v.(map[string]any)
		yr, _ := strconv.Atoi(data["Year"].(string))
		films = append(films, models.BasicFilm{
			ImdbId:      data["imdbID"].(string),
			Title:       data["Title"].(string),
			PosterUrl:   data["Poster"].(string),
			Type:        models.MovieType(data["Type"].(string)),
			ReleaseDate: time.Date(yr, 0, 0, 0, 0, 0, 0, time.Now().Location()),
		})
	}
	return films

}

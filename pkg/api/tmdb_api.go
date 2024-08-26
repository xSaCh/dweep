package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/xSaCh/dweep/pkg/models"
)

const (
	TmdbBaseUrl = "https://api.themoviedb.org/3/"
	ImageUrl    = "https://image.tmdb.org/t/p/"
)

type TmdbApi struct {
	apiKey string
}

func NewTmdbApi(apiKey string) *TmdbApi {
	return &TmdbApi{apiKey: apiKey}
}

// func (api *TmdbApi) SearchFilmByTitle(title string, seaechFilter *SearchFilter) []models.BasicFilm {
// 	print("Getting... ", TMDB_BASE_URL+"/search/movie?api_key="+api.apiKey+"&query="+title)
// 	res, err := http.Get(TMDB_BASE_URL + "/search/movie?api_key=" + api.apiKey + "&query=" + title)
// 	if err != nil {
// 		return []models.BasicFilm{}
// 	}
// 	defer res.Body.Close()

// 	data, _ := io.ReadAll(res.Body)
// 	var jsData map[string]interface{}
// 	json.Unmarshal(data, &jsData)

// 	// var films []models.BasicFilm
// 	// for _, v := range jsData["results"].([]interface{})  {
// 	// 	films = append(films, models.BasicFilm{
// 	// 		// ImdbId: v['i'],
// 	// 	})
// 	// }
// 	return []models.BasicFilm{}

// }

func (api *TmdbApi) GetMovie(tmdbId string) *models.Film {
	res, err := http.Get(TmdbBaseUrl + "movie/" + tmdbId + "?api_key=" + api.apiKey + "&append_to_response=external_ids,keywords")
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	var jsData map[string]interface{}
	json.Unmarshal(data, &jsData)

	imdbId := ""
	eIds := jsData["external_ids"].(map[string]interface{})
	if eIds["imdb_id"] != nil {
		imdbId = eIds["imdb_id"].(string)
	}

	genres := []string{}
	gM := jsData["genres"].([]interface{})
	for _, v := range gM {
		genres = append(genres, v.(map[string]interface{})["name"].(string))
	}

	keywords := []string{}
	kM := jsData["keywords"].(map[string]interface{})["keywords"].([]interface{})
	for _, v := range kM {
		keywords = append(keywords, v.(map[string]interface{})["name"].(string))
	}

	releaseDate, _ := time.Parse(time.DateOnly, jsData["release_date"].(string))

	return &models.Film{
		TmdbId:      int(jsData["id"].(float64)),
		ImdbId:      imdbId,
		Title:       jsData["original_title"].(string),
		Type:        models.Movie,
		Genres:      genres,
		ReleaseDate: releaseDate,
		Runtime:     int(jsData["runtime"].(float64)),
		Overview:    jsData["overview"].(string),
		PosterUrl:   jsData["poster_path"].(string),
		Rating:      float32(jsData["vote_average"].(float64)),
		Keywords:    keywords,

		Director:  "",
		MainCasts: []string{},
	}
}

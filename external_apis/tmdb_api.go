package externalapi

import (
	"encoding/json"
	"fmt"
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

func (api *TmdbApi) GetMovie(tmdbId string) *models.Movie {
	res, err := http.Get(TmdbBaseUrl + "movie/" + tmdbId + "?api_key=" + api.apiKey + "&append_to_response=external_ids,keywords,credits,release_dates")
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	var jsData map[string]interface{}
	json.Unmarshal(data, &jsData)

	if sB, ok := jsData["success"].(bool); ok && !sB {
		return nil
	}

	genres := []string{}
	gM := jsData["genres"].([]interface{})
	for _, v := range gM {
		genres = append(genres, v.(map[string]interface{})["name"].(string))
	}

	imdbId := ""
	if eIds, ok := jsData["external_ids"].(map[string]interface{}); ok {
		if eIds["imdb_id"] != nil {
			imdbId = eIds["imdb_id"].(string)
		}
	}

	keywords := []string{}
	if k, ok := jsData["keywords"].(map[string]interface{}); ok {
		kM := k["keywords"].([]interface{})
		for _, v := range kM {
			keywords = append(keywords, v.(map[string]interface{})["name"].(string))
		}
	}

	mainCasts := []string{}
	director := ""
	if cM, ok := jsData["credits"].(map[string]interface{}); ok {

		if casts, cOk := cM["cast"].([]interface{}); cOk {
			for i, v := range casts {
				if i == 4 {
					break
				}
				mainCasts = append(mainCasts, v.(map[string]interface{})["name"].(string))
			}
		}
		if crews, crOk := cM["crew"].([]interface{}); crOk {
			for _, v := range crews {
				if v.(map[string]interface{})["job"].(string) == "Director" {
					director = v.(map[string]interface{})["name"].(string)
					break
				}
			}
		}
	}

	ageRating := ""
	if cM, ok := jsData["release_dates"].(map[string]interface{}); ok {
		for _, v := range cM["results"].([]interface{}) {
			if v.(map[string]interface{})["iso_3166_1"].(string) == "US" {
				rM := v.(map[string]interface{})["release_dates"].([]interface{})
				for _, vr := range rM {
					if vr.(map[string]interface{})["type"].(float64) == 3 {
						ageRating = vr.(map[string]interface{})["certification"].(string)
						break
					}
				}
			}
		}
	}

	releaseDate, _ := time.Parse(time.DateOnly, jsData["release_date"].(string))

	return &models.Movie{
		Id:          -1,
		TmdbId:      int(jsData["id"].(float64)),
		ImdbId:      imdbId,
		Title:       jsData["title"].(string),
		Genres:      genres,
		ReleaseDate: releaseDate,
		Runtime:     int(jsData["runtime"].(float64)),
		AgeRating:   ageRating,

		Overview:  jsData["overview"].(string),
		PosterUrl: jsData["poster_path"].(string),
		Rating:    float32(jsData["vote_average"].(float64)),
		Tags:      keywords,

		Director:  director,
		MainCasts: mainCasts,
	}
}

func (api *TmdbApi) GetSeries(tmdbId string) *models.Show {
	res, err := http.Get(TmdbBaseUrl + "tv/" + tmdbId + "?api_key=" + api.apiKey + "&append_to_response=external_ids,keywords,credits,content_ratings")
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	var jsData map[string]interface{}
	json.Unmarshal(data, &jsData)

	if sB, ok := jsData["success"].(bool); ok && !sB {
		fmt.Println(jsData)
		return nil
	}

	genres := []string{}
	gM := jsData["genres"].([]interface{})
	for _, v := range gM {
		genres = append(genres, v.(map[string]interface{})["name"].(string))
	}

	imdbId := ""
	if eIds, ok := jsData["external_ids"].(map[string]interface{}); ok {
		if eIds["imdb_id"] != nil {
			imdbId = eIds["imdb_id"].(string)
		}
	}

	keywords := []string{}
	if k, ok := jsData["keywords"].(map[string]interface{}); ok {
		kM := k["results"].([]interface{})
		for _, v := range kM {
			keywords = append(keywords, v.(map[string]interface{})["name"].(string))
		}
	}

	mainCasts := []string{}
	director := ""
	if cM, ok := jsData["credits"].(map[string]interface{}); ok {

		if casts, cOk := cM["cast"].([]interface{}); cOk {
			for i, v := range casts {
				mainCasts = append(mainCasts, v.(map[string]interface{})["name"].(string))
				if i == 3 {
					break
				}
			}
		}
		if crews, crOk := cM["crew"].([]interface{}); crOk {
			for _, v := range crews {
				if v.(map[string]interface{})["job"].(string) == "Director" {
					director = v.(map[string]interface{})["name"].(string)
					break
				}
			}
		}
	}

	ageRating := ""
	if cM, ok := jsData["content_ratings"].(map[string]interface{}); ok {
		for _, v := range cM["results"].([]interface{}) {
			if v.(map[string]interface{})["iso_3166_1"].(string) == "US" {
				ageRating = v.(map[string]interface{})["rating"].(string)
			}
		}
	}

	firstDate, _ := time.Parse(time.DateOnly, jsData["first_air_date"].(string))
	lastDate, _ := time.Parse(time.DateOnly, jsData["last_air_date"].(string))

	_ = director
	return &models.Show{
		Id:        -1,
		TmdbId:    int(jsData["id"].(float64)),
		ImdbId:    imdbId,
		Title:     jsData["name"].(string),
		Genres:    genres,
		AgeRating: ageRating,
		Status:    models.SeriesStatus(jsData["status"].(string)),

		FirstAirDate: firstDate,
		LastAirDate:  lastDate,
		NoSeasons:    int(jsData["number_of_seasons"].(float64)),
		NoEpisodes:   int(jsData["number_of_episodes"].(float64)),

		Overview:  jsData["overview"].(string),
		PosterUrl: jsData["poster_path"].(string),
		Rating:    float32(jsData["vote_average"].(float64)),
		Keywords:  keywords,

		MainCasts: mainCasts,
	}
}

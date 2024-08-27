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
	res, err := http.Get(TmdbBaseUrl + "movie/" + tmdbId + "?api_key=" + api.apiKey + "&append_to_response=external_ids,keywords,credits,release_dates")
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	var jsData map[string]interface{}
	json.Unmarshal(data, &jsData)

	if jsData["success"] != nil && !jsData["success"].(bool) {
		return nil
	}

	genres := []string{}
	gM := jsData["genres"].([]interface{})
	for _, v := range gM {
		genres = append(genres, v.(map[string]interface{})["name"].(string))
	}

	imdbId := ""
	if jsData["external_ids"] != nil {
		eIds := jsData["external_ids"].(map[string]interface{})
		if eIds["imdb_id"] != nil {
			imdbId = eIds["imdb_id"].(string)
		}
	}

	keywords := []string{}
	if jsData["keywords"] != nil {
		kM := jsData["keywords"].(map[string]interface{})["keywords"].([]interface{})
		for _, v := range kM {
			keywords = append(keywords, v.(map[string]interface{})["name"].(string))
		}
	}
	TheatricalRelease := 3.0

	casts := []string{}
	director := ""
	if jsData["credits"] != nil {
		cM := jsData["credits"].(map[string]interface{})
		if cM["cast"] != nil {
			for i, v := range cM["cast"].([]interface{}) {

				casts = append(casts, v.(map[string]interface{})["name"].(string))
				if i == 3 {
					break
				}
			}
		}
		if cM["crew"] != nil {
			for _, v := range cM["crew"].([]interface{}) {
				if v.(map[string]interface{})["job"].(string) == "Director" {
					director = v.(map[string]interface{})["name"].(string)
					break
				}
			}
		}
	}

	ageRating := ""
	if jsData["release_dates"] != nil {
		cM := jsData["release_dates"].(map[string]interface{})
		for _, v := range cM["results"].([]interface{}) {
			if v.(map[string]interface{})["iso_3166_1"].(string) == "US" {
				rM := v.(map[string]interface{})["release_dates"].([]interface{})
				for _, vr := range rM {
					if vr.(map[string]interface{})["type"].(float64) == TheatricalRelease {
						ageRating = vr.(map[string]interface{})["certification"].(string)
						break
					}
				}
			}
		}
	}

	releaseDate, _ := time.Parse(time.DateOnly, jsData["release_date"].(string))

	return &models.Film{
		TmdbId:      int(jsData["id"].(float64)),
		ImdbId:      imdbId,
		Title:       jsData["title"].(string),
		Type:        models.Movie,
		Genres:      genres,
		ReleaseDate: releaseDate,
		Runtime:     int(jsData["runtime"].(float64)),
		AgeRating:   ageRating,

		Overview:  jsData["overview"].(string),
		PosterUrl: jsData["poster_path"].(string),
		Rating:    float32(jsData["vote_average"].(float64)),
		Keywords:  keywords,

		Director:  director,
		MainCasts: casts,
	}
}

func (api *TmdbApi) GetSeries(tmdbId string) *models.FilmSeries {
	res, err := http.Get(TmdbBaseUrl + "tv/" + tmdbId + "?api_key=" + api.apiKey + "&append_to_response=external_ids,keywords,credits,content_ratings")
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	var jsData map[string]interface{}
	json.Unmarshal(data, &jsData)

	if jsData["success"] != nil && !jsData["success"].(bool) {
		return nil
	}

	genres := []string{}
	gM := jsData["genres"].([]interface{})
	for _, v := range gM {
		genres = append(genres, v.(map[string]interface{})["name"].(string))
	}

	imdbId := ""
	if jsData["external_ids"] != nil {
		eIds := jsData["external_ids"].(map[string]interface{})
		if eIds["imdb_id"] != nil {
			imdbId = eIds["imdb_id"].(string)
		}
	}

	keywords := []string{}
	if jsData["keywords"] != nil {
		kM := jsData["keywords"].(map[string]interface{})["results"].([]interface{})
		for _, v := range kM {
			keywords = append(keywords, v.(map[string]interface{})["name"].(string))
		}
	}

	casts := []string{}
	director := ""
	if jsData["credits"] != nil {
		cM := jsData["credits"].(map[string]interface{})
		if cM["cast"] != nil {
			for i, v := range cM["cast"].([]interface{}) {

				casts = append(casts, v.(map[string]interface{})["name"].(string))
				if i == 3 {
					break
				}
			}
		}
		if cM["crew"] != nil {
			for _, v := range cM["crew"].([]interface{}) {
				if v.(map[string]interface{})["job"].(string) == "Director" {
					director = v.(map[string]interface{})["name"].(string)
					break
				}
			}
		}
	}

	ageRating := ""
	if jsData["content_ratings"] != nil {
		cM := jsData["content_ratings"].(map[string]interface{})
		for _, v := range cM["results"].([]interface{}) {
			if v.(map[string]interface{})["iso_3166_1"].(string) == "US" {
				ageRating = v.(map[string]interface{})["rating"].(string)
			}
		}
	}

	releaseDate, _ := time.Parse(time.DateOnly, jsData["first_air_date"].(string))

	return &models.FilmSeries{
		Film: models.Film{
			TmdbId:      int(jsData["id"].(float64)),
			ImdbId:      imdbId,
			Title:       jsData["name"].(string),
			Type:        models.Series,
			Genres:      genres,
			ReleaseDate: releaseDate,
			Runtime:     0,
			AgeRating:   ageRating,

			Overview:  jsData["overview"].(string),
			PosterUrl: jsData["poster_path"].(string),
			Rating:    float32(jsData["vote_average"].(float64)),
			Keywords:  keywords,

			Director:  director,
			MainCasts: casts},

		Status:        models.SeriesStatus(jsData["status"].(string)),
		TotalSeasons:  int(jsData["number_of_seasons"].(float64)),
		TotalEpisodes: int(jsData["number_of_episodes"].(float64)),
	}
}

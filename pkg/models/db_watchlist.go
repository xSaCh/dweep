package models

import "time"

type DBFilmWatchlistItem struct {
	FilmId int
	Type   FilmType

	MyRating    float32
	MyTags      []string
	WatchStatus WatchStatus

	Note          string
	RecommendedBy []int64
	WatchedDate   time.Time
}

type DBFilmWatchlistItemMovie struct {
	DBFilmWatchlistItem
	WatchedDates []time.Time
}

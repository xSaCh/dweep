package models

import "time"

type ReqWatchlistItem struct {
	Id int

	MyRating    float32
	MyTags      []string
	WatchStatus WatchStatus

	Note          string
	RecommendedBy []int64
}

type ReqWatchlistItemMovie struct {
	ReqWatchlistItem
	WatchedDates []time.Time
}

// type ReqWatchlistItemShow struct {
// 	ReqWatchlistItem
// 	WatchedDates []time.Time
// }

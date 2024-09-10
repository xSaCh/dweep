package models

import "time"

type ReqWatchlistItem struct {
	MyRating    float32     `json:"myRating"`
	MyTags      []string    `json:"myTags"`
	WatchStatus WatchStatus `json:"watchStatus"`

	Note          string  `json:"note"`
	RecommendedBy []int64 `json:"recommendedBy"`
}

type ReqWatchlistItemMovie struct {
	ReqWatchlistItem
	WatchedDates []time.Time `json:"watchDates"`
}

// type ReqWatchlistItemShow struct {
// 	ReqWatchlistItem
// 	WatchedDates []time.Time
// }

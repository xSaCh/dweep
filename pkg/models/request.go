package models

import "time"

type ReqWatchlistItem struct {
	MyRating    float32     `json:"myRating,omitempty"`
	MyTags      []string    `json:"myTags,omitempty"`
	WatchStatus WatchStatus `json:"watchStatus"`

	Note          string  `json:"note,omitempty"`
	RecommendedBy []int64 `json:"recommendedBy,omitempty"`
}

type ReqWatchlistItemMovie struct {
	ReqWatchlistItem
	WatchedDates []time.Time `json:"watchDates"`
}

// type ReqWatchlistItemShow struct {
// 	ReqWatchlistItem
// 	WatchedDates []time.Time
// }

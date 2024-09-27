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

type ReqWatchlistItemShow struct {
	ReqWatchlistItem
	Episodes []ReqEpItem `json:"episodes"`
}

type ReqEpItem struct {
	EpisodeId   int64     `json:"episodeId"`
	WatchedDate time.Time `json:"watchedDate"`
}

// type ReqWatchlistItemShow struct {
// 	ReqWatchlistItem
// 	WatchedDates []time.Time
// }

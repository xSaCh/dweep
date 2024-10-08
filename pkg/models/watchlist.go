package models

import "time"

type WhenToWatch string
type WatchStatus string

const (
	Someday     WhenToWatch = "someday"
	WithFocused WhenToWatch = "with_focused"
	WithFamily  WhenToWatch = "with_family"
	WithFriend  WhenToWatch = "with_friend"
	Sad         WhenToWatch = "sad"
	Happy       WhenToWatch = "happy"
	Frustrate   WhenToWatch = "frustrate"
	Motivate    WhenToWatch = "motivate"

	Watched     WatchStatus = "watched"
	PlanToWatch WatchStatus = "plan_to_watch"
	Dropped     WatchStatus = "dropped"
	Watching    WatchStatus = "watching" // Generally for series
	OnHold      WatchStatus = "on_hold"  // Generally for series
)

type WatchlistItem struct {
	WatchlistItemId int      `json:"watchlist_item_id"`
	FilmId          int      `json:"film_id"`
	Type            FilmType `json:"type"`

	MyRating    float32     `json:"my_rating"` // 0 means not rated
	MyTags      []string    `json:"my_tags"`
	WatchStatus WatchStatus `json:"watch_status"`

	Note          string  `json:"note"`
	RecommendedBy []int64 `json:"recommended_by"`

	AddedOn   time.Time `json:"added_on"`
	UpdatedOn time.Time `json:"updated_on"`
}
type WatchlistItemMovie struct {
	WatchlistItem
	WatchedDates []time.Time `json:"watched_dates"`
	// TimedWatched
}
type WatchlistItemShow struct {
	WatchlistItem
	WatchedEpisodes []WatchedEpisode `json:"watched_episodes"`
}

type WatchedEpisode struct {
	EpisodeId   int64     `json:"episode_id"`
	WatchedDate time.Time `json:"watched_date"`
}

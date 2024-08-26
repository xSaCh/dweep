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

type Watchlist struct {
	Film
	MyRating    float32     `json:"my_rating"`
	MyTags      []string    `json:"my_tags"`
	WatchStatus WatchStatus `json:"watch_status"`

	Note          string      `json:"note"`
	RecommendedBy []string    `json:"recommended_by"`
	WhenToWatch   WhenToWatch `json:"when_to_watch"`
}
type WatchlistMovie struct {
	Watchlist
	WatchedDates []time.Time `json:"watched_dates"`
	// TimedWatched
}
type WatchedEpisode struct {
	EpisodeId   int64     `json:"episode_id"`
	WatchedDate time.Time `json:"watched_date"`
}
type WatchlistSeries struct {
	Watchlist
	FullyWatchedSeasons []int            `json:"fully_watched_seasons"`
	WatchedEpisodes     []WatchedEpisode `json:"watched_episodes"`
}

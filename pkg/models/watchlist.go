package models

/*
Watchlist
    Note
    Recommended by
    When to watch
        someday
        when sad
        with focused
        with family
        with friend
*/

type Watchlist struct {
	MovId      int    `gorm:"primaryKey" json:"movId"`
	TmdbId     string `json:"tmdbId" gorm:"primaryKey;autoIncrement:false"`
	AddedOn    string `json:"addedOn"`
	Watched    bool   `json:"watched"`
	WatchCount int    `json:"watchCount,omitempty"`
}

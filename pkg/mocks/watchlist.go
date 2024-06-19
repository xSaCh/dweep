package mocks

import (
	"time"

	"github.com/xSaCh/dweep/pkg/models"
)

/*
393    tt4160708   yes      2023-09-08
392    tt10954652
391    tt9844522   yes      2023-09-08
390    tt6723592   yes      2023-09-07
389    tt7601480
388    tt4139588
387    tt8130968
*/

func addDate(date string) *time.Time {
	a, _ := time.Parse(time.DateOnly, date)
	return &a
	// return date
}

var Watchlists = []models.Watchlist{
	{MovId: 389, TmdbId: "tt7601480", Watched: false, WatchCount: 0, AddedOn: "2023-09-07"},
	{MovId: 390, TmdbId: "tt6723592", Watched: true, WatchCount: 1, AddedOn: "2023-09-07"},
	{MovId: 391, TmdbId: "tt9844522", Watched: true, WatchCount: 1, AddedOn: "2023-09-08"},
	{MovId: 392, TmdbId: "tt10954652", Watched: false, WatchCount: 0, AddedOn: "2023-09-07"},
	{MovId: 393, TmdbId: "tt4160708", Watched: true, WatchCount: 1, AddedOn: "2023-09-08"},
}

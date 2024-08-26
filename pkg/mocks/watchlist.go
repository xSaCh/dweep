package mocks

import (
	"time"
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
}

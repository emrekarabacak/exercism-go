package gigasecond

import (
	"time"
)

const oneGs = 1e9 * time.Second

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(oneGs)
}

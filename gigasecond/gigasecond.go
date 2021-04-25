package gigasecond

import (
	"time"
)

// AddGigasecond add 10^9 to the time and return the new time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}

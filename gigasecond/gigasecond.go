package gigasecond

import "time"

// AddGigasecond returns a time that's 10^9 seconds later than the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}

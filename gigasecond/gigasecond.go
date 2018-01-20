package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond Calculate the moment when someone has lived for 10^9 seconds.
func AddGigasecond(t time.Time) time.Time {
	aGigaSecond := time.Duration(math.Pow10(9))
	return t.Add(time.Second * aGigaSecond)
}

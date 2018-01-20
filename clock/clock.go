package clock

import (
	"fmt"
	"time"
)

type Clock struct {
	hours   int
	minutes int
}

func New(hours int, minutes int) Clock {
	aDate := time.Date(0, 0, 0, hours, minutes, 0, 0, time.UTC)

	return Clock{aDate.Hour(), aDate.Minute()}
}

func (clock Clock) Add(minutes int) Clock {
	totalMinutes := clock.minutes + minutes
	aDate := time.Date(0, 0, 0, clock.hours, totalMinutes, 0, 0, time.UTC)

	return Clock{aDate.Hour(), aDate.Minute()}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hours, clock.minutes)
}

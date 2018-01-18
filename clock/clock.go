package clock

import (
	"fmt"
)

// Clock stores hour and minutes
type Clock struct {
	hours   int
	minutes int
}

// New is the constructor
func New(aHours int, aMinutes int) Clock {
	minutes := aMinutes % 60
	if minutes < 0 {
		minutes = 60 + minutes
		aHours--
	}
	hours := (aHours + (aMinutes / 60)) % 24
	if hours < 0 {
		hours = 24 + hours
	}
	return Clock{hours: hours, minutes: minutes}
}

// String returns the clock hour formatted with hh:mm
func (aClock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", aClock.hours, aClock.minutes)
}

// Add minutes to Clock
func (aClock Clock) Add(minutes int) Clock {
	return New(aClock.hours, minutes+aClock.minutes)
}

// Package gigasecond provides a function to calculate the time after adding
// one gigasecond.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond calculates the time after adding one gigasecond to the given
// time.
//
// It takes a time.Time parameter and returns a time.Time value.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}

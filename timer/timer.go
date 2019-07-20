package timer

import (
	"time"
)

// Actor
type Actor interface {
	Act(tag string, d time.Duration)
}

// Start returns a function that should be called at the end of the period to
// be timed (typically through use of the defer mechanism)
func Start(tag string, a Actor) func() {
	start := time.Now()

	return func() {
		d := time.Since(start)
		a.Act(tag, d)
	}
}

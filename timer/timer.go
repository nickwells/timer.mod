package timer

import (
	"time"
)

// Actor is the interface that must be satisfied by the second parameter to
// the Start function.
type Actor interface {
	Act(tag string, d time.Duration)
}

// Start returns a function that should be called at the end of the period to
// be timed (typically through use of the defer mechanism). When you call the
// returned function it will calculate the duration since Start was called
// and pass that with the tag to the Act method of the supplied Actor.
func Start(tag string, a Actor) func() {
	start := time.Now()

	return func() {
		d := time.Since(start)
		a.Act(tag, d)
	}
}

// StartF returns a function that should be called at the end of the period
// to be timed (typically through use of the defer mechanism). When you call
// the returned function it will calculate the duration since StartF was
// called and pass that with the tag to the supplied function, f
func StartF(tag string, f func(string, time.Duration)) func() {
	start := time.Now()

	return func() {
		d := time.Since(start)
		f(tag, d)
	}
}

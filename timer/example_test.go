package timer_test

import (
	"time"

	"github.com/nickwells/timer.mod/timer"
)

var secondPrinter = timer.Print{
	Fmt:   "%s: %d second\n",
	Units: time.Second,
}

func ExampleStart() {
	defer timer.Start("ExampleStart", secondPrinter)()
	time.Sleep(1 * time.Second)
	// Output:
	// ExampleStart: 1 second
}

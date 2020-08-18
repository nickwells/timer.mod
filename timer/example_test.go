package timer_test

import (
	"fmt"
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

func ExampleStartF() {
	defer timer.StartF("ExampleStartF",
		func(tag string, d time.Duration) {
			fmt.Printf("%s: %d second\n", tag, d/time.Second)
		})()
	time.Sleep(1 * time.Second)
	// Output:
	// ExampleStartF: 1 second
}

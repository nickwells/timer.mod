package timer

import (
	"fmt"
	"io"
	"os"
	"time"
)

// DefaultFmt is the default format used to print the tag and duration when
// the action function is called. To override this value you should set the
// Fmt member of the Print struct. Note that the parameters it must take are
// a string and a time.Duration, in that order.
const DefaultFmt = "%s: %d\n"

// Print implements the Actor interface and so is suitable for passing to
// the Start function. Its Act method will simply print the time prefixed
// with the tag
type Print struct {
	// Fmt can be set to override the default format. It must take a string
	// as the first and an integer as the second parameter
	Fmt string
	// Units can be set to replace the default duration units of Nanoseconds
	Units time.Duration
	// W can be set to override the default output (os.Stdout)
	W io.Writer
}

// Act prints the duration prefixed with the tag
func (p Print) Act(tag string, d time.Duration) {
	units := time.Nanosecond
	if p.Units > 0 {
		units = p.Units
	}

	f := DefaultFmt
	if p.Fmt != "" {
		f = p.Fmt
	}

	var w io.Writer = os.Stdout
	if p.W != nil {
		w = p.W
	}

	fmt.Fprintf(w, f, tag, d/units)
}

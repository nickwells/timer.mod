/*
Package timer provides a means of measuring the duration of a piece of code.

You call the Start function passing it a string describing the task being
timed and an actor to be called on completion. Then you simply call the
returned function when the task completes.

There is a sample Actor provided called Print which will simply print the tag
and duration.
*/
package timer

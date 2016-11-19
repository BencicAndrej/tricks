package clock

import "time"

// Clock wraps the `Now()` function from the `time` package.
//
// By hiding the `time.Now()` behind an interface, we explicitly
// declare time as a dependency, and we enable the option of
// providing fake Clock implementations for testing purposes.
type Clock interface {
	Now() time.Time
}

// New function returns a new Clock implementation.
func New() Clock {
	return &realClock{}
}

// realClock simply wraps the `time.Now()` method.
type realClock struct{}

// Now method satisfies the `Clock` interface using the `Now`
// method from the `time` package.
func (*realClock) Now() time.Time {
	return time.Now()
}

// BrokenClock is a mock implementation of the `Clock` interface
// that always returns the same time.
type BrokenClock struct {
	t time.Time
}

// NewBrokenClock returns a broken clock that always returns the
// time provided at initialization.
func NewBrokenClock(t time.Time) *BrokenClock {
	return &BrokenClock{
		t: t,
	}
}

// Now always returns time provided at initialization.
func (c *BrokenClock) Now() time.Time {
	return c.t
}

// SetTime replaces the internal time that always gets returned
// by `BrokenClock.Now()`
func (c *BrokenClock) SetTime(t time.Time) {
	c.t = t
}

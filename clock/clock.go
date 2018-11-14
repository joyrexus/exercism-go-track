package clock

import "fmt"

// A Clock represents an instant in the day with minute precision.
type Clock int

// Minutes in an hour.
const H int = 60

// Minutes in a day.
const D int = 24 * H

// New is a constructor for a clock.
func New(h, m int) Clock {
	m = (H*h + m) % D
	if m < 0 {
		m += D
	}
	return Clock(m)
}

// String is a "stringer" method the clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add adds minutes to the clock.
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract subtracts minutes from the clock.
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

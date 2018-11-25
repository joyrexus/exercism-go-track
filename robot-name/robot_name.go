package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot represents a robot with a name.
type Robot struct {
	name string
}

// seen is used to track previously seen robot names.
var seen = map[string]bool{"": true}

// Name returns the robot's name.
func (r *Robot) Name() string {
	if r.name == "" {
		r.Reset()
	}
	return r.name
}

// Reset resets the robot's name.
func (r *Robot) Reset() {
	rand.Seed(time.Now().UnixNano())
	for seen[r.name] {
		r.name = fmt.Sprintf("%c%c%03d",
			rand.Intn(26)+'A',
			rand.Intn(26)+'A',
			rand.Intn(1000))
	}
	seen[r.name] = true
}

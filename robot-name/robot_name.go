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

// randLetter returns a random rune in the range A-Z.
func randLetter() int {
	return rand.Intn(26) + 'A'
}

// makeLetters makes a random two-character string of uppercase letters.
func makeLetters() string {
	return fmt.Sprintf("%c%c", randLetter(), randLetter())
}

// makeNumbers makes a random three-digit string.
func makeNumbers() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

// generate generates a random robot name of the form AA###.
func (r *Robot) generate() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%s", makeLetters(), makeNumbers())
}

// Name returns the robot's name.
func (r *Robot) Name() string {
	if r.name == "" {
		r.Reset()
	}
	return r.name
}

// Reset resets the robot's name.
func (r *Robot) Reset() {
	for seen[r.name] {
		r.name = r.generate()
	}
	seen[r.name] = true
}

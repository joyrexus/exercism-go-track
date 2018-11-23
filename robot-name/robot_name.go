package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot represents a robot.
type Robot struct {
	name string
}

var seen = map[string]bool{"": true}

var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randLetter() string {
	return string(letters[rand.Intn(len(letters))])
}

func makeLetters() string {
	return fmt.Sprintf("%s%s", randLetter(), randLetter())
}

func makeNumbers() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

// generate generates a random robot name.
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

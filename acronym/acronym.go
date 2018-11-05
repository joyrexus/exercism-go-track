package acronym

import (
	"regexp"
	"strings"
)

var nonword = regexp.MustCompile(`\W+`)

// Abbreviate returns the acronym form of the given input string.
func Abbreviate(s string) (a string) {
	for _, w := range nonword.Split(s, -1) {
		if len(w) > 1 {
			a += string(w[0])
		}
	}
	return strings.ToUpper(a)
}

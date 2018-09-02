package acronym

import (
	"regexp"
	"strings"
)

var whitespace = regexp.MustCompile(`\W+`)

// Abbreviate returns the acronym form of the given input string.
func Abbreviate(s string) string {
	var a []string
	for _, w := range whitespace.Split(s, -1) {
		if len(w) > 1 {
			a = append(a, strings.ToUpper(string(w[0])))
		}
	}
	return strings.Join(a, "")
}

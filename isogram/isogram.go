package isogram

import "strings"

// IsIsogram tests whether a string is an isogram.
func IsIsogram(s string) bool {
	var chars = map[rune]int{}
	for _, r := range strings.ToLower(s) {
		if strings.Contains(" -", string(r)) {
			continue
		}
		if chars[r] == 1 {
			return false
		}
		chars[r] = 1
	}
	return true
}

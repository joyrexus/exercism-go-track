// Package bob contains utility functions for chatting with Bob.
package bob

import (
	"regexp"
	"strings"
)

var HasCaps = regexp.MustCompile(`[A-Z]`).MatchString
var IsSilent = regexp.MustCompile(`^\s*$`).MatchString

// Hey returns Bob's responses which vary based on the type of remark.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case IsSilent(remark):
		return "Fine. Be that way!"
	case HasCaps(remark) && strings.ToUpper(remark) == remark:
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}

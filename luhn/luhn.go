package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if a string is a kosher Luhn value.
func Valid(v string) bool {
	v = strings.Replace(v, " ", "", -1)

	// check length and confirm we only have digits
	if len(v) < 2 {
		return false
	}
	if _, err := strconv.Atoi(v); err != nil {
		return false
	}

	sum := 0		// sum of digits in value after doubling
	double := false // every other digit from right

	for i := len(v) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(string(v[i]))
		if double {
			x = 2 * x
			if x > 9 {
				x = x - 9
			}
		}
		double = false == double // toggle value
		sum += x
	}
	return sum%10 == 0
}

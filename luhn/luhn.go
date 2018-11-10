package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if a string is a kosher Luhn value.
func Valid(v string) bool {
	v = strings.Replace(v, " ", "", -1)
	if len(v) < 2 {
		return false
	}

	checksum := 0           // the sum after doubling
	double := len(v)%2 == 0 // every other digit from right

	// iterate over each digit in value to get checksum
	for i := range v {
		x, err := strconv.Atoi(string(v[i]))
		if err != nil {
			return false
		}
		if double {
			x = 2 * x
			if x > 9 {
				x = x - 9
			}
		}
		double = !double // toggle value
		checksum += x
	}
	return checksum%10 == 0
}

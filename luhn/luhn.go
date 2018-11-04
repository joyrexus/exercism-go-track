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

	checksum := 0   // sum of digits in value after doubling
	double := false // every other digit from right

	// iterate over each digit in value to get checksum
	for i := len(v) - 1; i >= 0; i-- {
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
		double = false == double // toggle value
		checksum += x
	}
	return checksum%10 == 0
}

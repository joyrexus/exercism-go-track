package raindrops

import "strconv"

// Convert takes an integer and returns a string that varies based on the
// number's factorization.
func Convert(i int) string {
	var r string
	if i%3 == 0 {
		r = "Pling"
	}
	if i%5 == 0 {
		r += "Plang"
	}
	if i%7 == 0 {
		r += "Plong"
	}
	if r == "" {
		r = strconv.Itoa(i)
	}
	return r
}

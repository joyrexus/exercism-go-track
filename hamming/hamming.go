package hamming

import "errors"

// Distance returns the hamming distance between two strings.
func Distance(a, b string) (int, error) {
	switch {
	case a == b:
		return 0, nil
	case len(a) != len(b):
		return -1, errors.New("string lengths differ")
	default:
		d := 0
		for i := range a {
			if a[i] != b[i] {
				d++
			}
		}
		return d, nil
	}
}

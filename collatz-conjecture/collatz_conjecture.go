package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1 when
// applying the collatz algorithm on the given input.
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("input must be a positive integer")
	}

	var steps int
	for n > 1 {
		steps++
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
	}
	return steps, nil
}

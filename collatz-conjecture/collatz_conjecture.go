package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1 when
// applying the collatz algorithm on the given input.
func CollatzConjecture(n int) (int, error) {
	steps := 0
	for {
		if n < 1 {
			return steps, errors.New("input must be a positive integer")
		}
		if n == 1 {
			return steps, nil
		}

		steps++

		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
	}
}

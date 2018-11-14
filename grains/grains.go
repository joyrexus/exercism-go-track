package grains

import "errors"

// Square returns the number of grains on the nth square.
func Square(n int) (uint64, error) {
	if 1 > n || n > 64 {
		return 1, errors.New("Invalid input")
	}
	return 1 << uint(n-1), nil
}

// Total returns the total number of grains on all 64 squares.
func Total() uint64 {
	return 1<<64 - 1
}

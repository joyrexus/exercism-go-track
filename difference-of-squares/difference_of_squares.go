package diffsquares

import "math"

// SquareOfSums returns the square of the sums of the first N natural numbers.
func SquareOfSums(n int) int {
	var sum float64
	for i := 1; i <= n; i++ {
		sum += float64(i)
	}
	return int(math.Pow(sum, 2))
}

// SumOfSquares returns the sum of the squares of the first N natural numbers.
func SumOfSquares(n int) int {
	var sum float64
	for i := 1; i <= n; i++ {
		sum += math.Pow(float64(i), 2)
	}
	return int(sum)
}

// Difference returns the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

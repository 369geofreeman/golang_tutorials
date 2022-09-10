package diffsquares

import "math"

// The square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	total := 0
	for x := 1; x < n+1; x++ {
		total += x
	}
	return int(math.Pow(float64(total), 2))
}

// The sum of the squares of the first N natural numbers
func SumOfSquares(n int) int {
	total := 0
	for x := 1; x < n+1; x++ {
		total += int(math.Pow(float64(x), 2))
	}
	return total
}

// The difference between the square of the sum of the first N natural numbers
// and the sum of the squares of the first N natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

package diffsquares

// SquareOfSum return square of first n numbers
func SquareOfSum(n uint64) uint64 {
	return ((n * (n + 1)) * (n * (n + 1))) / 4
}

// SumOfSquares return sum of first n numbers square
func SumOfSquares(n uint64) uint64 {
	return (n * (n + 1) * ((2 * n) + 1)) / 6
}

// Difference returns difference between SquareOfSum and SumOfSquares
func Difference(n uint64) uint64 {
	return SquareOfSum(n) - SumOfSquares(n)
}

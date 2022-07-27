package diffsquares

func SquareOfSum(n int) int {
	i := 1
	var count int
	for i <= n {
		count += i
		i++
	}
	return count * count
}

func SumOfSquares(n int) int {
	i := 1
	var count int
	for i <= n {
		count += i * i
		i++
	}
	return count
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

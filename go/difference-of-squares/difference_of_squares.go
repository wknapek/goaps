package diffsquares

func SquareOfSum(num int) int {
	var sum = 0
	for idx := 0; idx <= num; idx++ {
		sum += idx
	}
	return sum * sum
}

func SumOfSquares(num int) int {
	var sum = 0
	for idx := 0; idx <= num; idx++ {
		sum += idx * idx
	}
	return sum
}

func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}

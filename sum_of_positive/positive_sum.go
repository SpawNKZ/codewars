package kata

func PositiveSum(numbers []int) int {
	var result int

	for _, v := range numbers {
		if v > 0 {
			result += v
		}
	}

	return result
}

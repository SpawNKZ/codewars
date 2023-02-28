package multiples_of_3_or_5

func Multiple3And5(number int) int {
	var res int
	for i := 0; i < number; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			res += i
		}
	}
	return res
}

package recursion

func SumNumbers(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	if len(numbers) < 2 {
		return numbers[0]
	}

	number := numbers[0]

	return number + SumNumbers(numbers[1:])
}

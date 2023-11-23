package quicksort

func OrderNumbers(numbers []int) []int {
	length := len(numbers)
	if length < 2 {
		return numbers
	}

	if length == 2 && numbers[0] < numbers[1] {
		return numbers
	}

	if length == 2 && numbers[0] > numbers[1] {
		return []int{numbers[1], numbers[0]}
	}

	mid := length / 2
	number := numbers[mid]

	minors := []int{}
	bigger := []int{}

	rest := append(numbers[:mid], numbers[mid+1:]...)

	for i := 0; i < len(rest); i++ {
		if rest[i] < number {
			minors = append(minors, rest[i])
		} else {
			bigger = append(bigger, rest[i])
		}
	}

	return append(append(OrderNumbers(minors), number), OrderNumbers(bigger)...)
}

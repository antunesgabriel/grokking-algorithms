package selectionsort

func getSmallest(numbers []int) (int, int) {
	smallestNumber := numbers[0]
	smallestNumberIndex := 0

	for idx, num := range numbers {
		if num < smallestNumber {
			smallestNumber = num
			smallestNumberIndex = idx
		}
	}

	return smallestNumber, smallestNumberIndex
}

func removeNumber(numbers []int, idx int) []int {
	return append(numbers[:idx], numbers[idx+1:]...)
}

func SelectionSortNumbers(numbers []int) []int {
	numbersSorted := make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		smallest, idx := getSmallest(numbers)
		numbersSorted[i] = smallest

		numbers = removeNumber(numbers, idx)
	}

	return numbersSorted
}

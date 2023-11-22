package selectionsort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func createRandomSlice(length int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	numbers := make([]int, length)

	for i := 0; i < length; i++ {
		numbers[i] = r.Intn(100)
	}

	return numbers
}

func TestSelectionSortNumbers(t *testing.T) {
	t.Run("should sort numbers", func(t *testing.T) {
		numbers := []int{5, 6, 9, 10, 45, 3, 1, 6}

		expected := []int{1, 3, 5, 6, 6, 9, 10, 45}

		result := SelectionSortNumbers(numbers)

		if reflect.DeepEqual(result, expected) {
			t.Errorf("expected: %v - received %v\n", expected, result)
		}
	})
}

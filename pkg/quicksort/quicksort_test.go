package quicksort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestOrderNumbers(t *testing.T) {
	t.Run("OrderNumbers() should return slice if slice have length < 2", func(t *testing.T) {
		numbers := []int{9}

		result := OrderNumbers(numbers)

		if !reflect.DeepEqual(numbers, result) {
			t.Errorf("expect %v - receive %v\n", numbers, result)
		}
	})

	t.Run("OrderNumbers() should sort two item in slice", func(t *testing.T) {
		numbers := []int{55, 42}
		expected := []int{42, 55}

		result := OrderNumbers(numbers)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("expect %v - receive %v\n", expected, result)
		}
	})

	t.Run("OrderNumber() should sort big slice", func(t *testing.T) {
		numbers := []int{100, 45, 33, 66, 77, 45, 14, 7, 42}
		expected := []int{7, 14, 33, 42, 45, 45, 66, 77, 100}

		result := OrderNumbers(numbers)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expect %v - receive %v\n", expected, result)
		}
	})
}

func createNumbers(length int) []int {
	numbers := make([]int, length)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < length; i++ {
		numbers[i] = r.Intn(100)
	}

	return numbers
}

func BenchmarkOrderNumber(b *testing.B) {
	numbers := createNumbers(100)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = OrderNumbers(numbers)
	}
}

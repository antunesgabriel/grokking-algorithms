package recursion

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func createNumbers(length int) ([]int, int) {
	numbers := make([]int, length)
	sum := 0

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < length; i++ {
		numbers[i] = r.Intn(100)
		sum = sum + numbers[i]
	}

	return numbers, sum
}

func TestSumNumbers(t *testing.T) {
	t.Run("SumNumbers() should return 0 if slice is empty", func(t *testing.T) {
		result := SumNumbers([]int{})

		if result != 0 {
			t.Errorf("expect %d - receive %d", 0, result)
		}
	})

	t.Run("SumNumbers() should sum all numbers in slice", func(t *testing.T) {
		numbers, total := createNumbers(10)

		fmt.Printf("sum number: %v\n", numbers)

		result := SumNumbers(numbers)

		if result != total {
			t.Errorf("expect %d - receive %d - numbers %v", total, result, numbers)
		}
	})
}

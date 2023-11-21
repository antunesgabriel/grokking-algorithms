package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestFindNameBinarySearch(t *testing.T) {
	NAMES := []string{
		"Amanda",
		"Bruno",
		"Camila",
		"Diego",
		"Eduarda",
		"Fábio",
		"Giovana",
		"Henrique",
		"Isadora",
		"Júlio",
		"Larissa",
		"Matheus",
		"Natália",
		"Otávio",
		"Patrícia",
		"Ricardo",
		"Sabrina",
		"Thiago",
		"Ursula",
		"Vinícius",
	}

	t.Run("should return -1 if name not exist", func(t *testing.T) {
		index := FindNameBinarySearch(NAMES, "Unknow")

		if index != -1 {
			t.Errorf("Expected: %d - Received: %d", -1, index)
		}
	})

	t.Run("should return correct index to random name in slice", func(t *testing.T) {
		rand.Seed(time.Now().UnixNano())

		randomIndex := rand.Intn(len(NAMES))

		name := NAMES[randomIndex]

		result := FindNameBinarySearch(NAMES, name)

		if result != randomIndex {
			t.Errorf("Expected: %d - Received: %d", randomIndex, result)
		}
	})
}

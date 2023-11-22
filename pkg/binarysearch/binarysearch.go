package binarysearch

import (
	"strings"
)

func FindNameBinarySearch(names []string, name string) int {
	finish := len(names) - 1

	for start := 0; start <= finish; {
		meio := (start + finish) / 2
		item := names[meio]

		result := strings.Compare(item, name)

		if result == 0 {
			return meio
		}

		if result == -1 {
			start = meio + 1

			continue
		}

		finish = meio - 1
	}

	return -1
}

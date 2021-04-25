// Package hamming provides method to calculate hamming distance
package hamming

import (
	"errors"
)

// Distance returns hamming distance between two strings
func Distance(a, b string) (int, error) {
	lenA := len(a)
	lenB := len(b)

	if lenA != lenB {
		return 0, errors.New("Strings are not equal")
	}

	distance := 0

	for i := 0; i < lenA; i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}

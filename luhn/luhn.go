package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if the given string is valid as per luhn algorithm
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	iLen := len(input)

	if iLen <= 1 {
		return false
	}

	sum := 0
	isAlternate := iLen%2 == 0

	for _, r := range input {
		num, err := strconv.Atoi(string(r))

		if err != nil {
			return false
		}

		if isAlternate {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		sum += num
		isAlternate = !isAlternate
	}

	return (sum % 10) == 0
}

package isogram

import "unicode"

// IsIsogram return if input string is isogram or not
func IsIsogram(phrase string) bool {
	freq := make(map[rune]bool)

	for _, ch := range phrase {
		if ch == '-' || ch == ' ' {
			continue
		}

		upperChar := unicode.ToUpper(ch)

		if freq[upperChar] {
			return false
		}

		freq[upperChar] = true
	}

	return true
}

package raindrops

import "strconv"

// Convert return the raindrop sounds for the factors of the number
func Convert(raindropNumber int) string {
	output := ""

	if (raindropNumber % 3) == 0 {
		output += "Pling"
	}

	if (raindropNumber % 5) == 0 {
		output += "Plang"
	}

	if (raindropNumber % 7) == 0 {
		output += "Plong"
	}

	if output == "" {
		return strconv.Itoa(raindropNumber)
	}

	return output
}

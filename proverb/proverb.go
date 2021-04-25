package proverb

// Proverb returns the proverb based on the input.
func Proverb(rhyme []string) []string {
	i := 0
	rLen := len(rhyme)

	var output = make([]string, rLen)

	for i < rLen-1 {
		output[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
		i++
	}

	if rLen > 0 {
		output[i] = "And all for the want of a " + rhyme[0] + "."
	}

	return output
}

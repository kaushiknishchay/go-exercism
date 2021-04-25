package strand

import "strings"

// ToRNA converts DNA to RNA
func ToRNA(dna string) string {
	// output := ""

	// for _, ascii := range dna {
	// 	char := string(ascii)

	// 	switch char {
	// 	case "G":
	// 		output += "C"
	// 	case "C":
	// 		output += "G"
	// 	case "T":
	// 		output += "A"
	// 	case "A":
	// 		output += "U"
	// 	}
	// }

	// return output

	// return strings.NewReplacer(
	// 	"G", "C",
	// 	"C", "G",
	// 	"T", "A",
	// 	"A", "U",
	// ).Replace(dna)

	return strings.Map(func(r rune) rune {
		switch r {
		case 'G':
			return 'C'
		case 'C':
			return 'G'
		case 'T':
			return 'A'
		case 'A':
			return 'U'
		}
		return r
	}, dna)
}

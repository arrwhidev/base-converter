package radix

import "strings"

var (
	Binary      = []string{"0", "1"}
	Quaternary  = []string{"0", "1", "2", "3"}
	Octal       = []string{"0", "1", "2", "3", "4", "5", "6", "7"}
	Hexadecimal = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	Base62      = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

// Encode converts a base10 number to any alphabet.
func Encode(n uint64, alphabet []string) string {
	if n == 0 {
		return string(alphabet[0])
	}
	radix := uint64(len(alphabet))
	encoded := ""
	for ; n > 0; n = n / radix {
		encoded = string(alphabet[n%radix]) + encoded
	}
	return encoded
}

// Decode converts an encoded value to base10.
func Decode(input string, alphabet []string) uint64 {
	if input == string(alphabet[0]) {
		return 0
	} else if input == "" {
		panic("Input must not be empty.")
	}

	radix := uint64(len(alphabet))
	decoded := uint64(0)
	for _, c := range strings.Split(input, "") {
		alphabetIndex := uint64(indexOf(alphabet, c))
		decoded = radix*decoded + alphabetIndex
	}
	return decoded
}

func indexOf(a []string, s string) int {
	for i, el := range a {
		if s == el {
			return i
		}
	}
	panic("Alphabet out of bounds")
}

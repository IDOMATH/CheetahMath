package base_change

import (
	"github.com/IDOMATH/CheetahMath/int_pow"
	"strings"
)

const runeDifference = 87

func ToTen(base int, in string) int64 {
	in = strings.ToLower(in)
	var baseTen int64 = 0
	for i := 0; i < len(in); i++ {
		char := rune(in[len(in)-1-i])
		if char-'0' > 9 {
			char = char - runeDifference
		} else {
			char = char - '0'
		}
		baseTen = baseTen + (int64(char) * int_pow.IntPow(int64(base), int64(i)))
	}
	return baseTen
}

// FromTen takes an int base value, up to 36 and returns a string.
// If the base being converted to is 10 or less, the string should be
// able to have strings.Atoi() called directly on it for an integer value.
// Will return the empty string if an error occurs
func FromTen(base int, in int64) string {

	return ""
}

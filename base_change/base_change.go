package base_change

import (
	"strconv"
	"strings"

	"github.com/IDOMATH/CheetahMath/int_pow"
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
func FromTen(base int64, in int64) string {
	var reversedResult strings.Builder
	if in < 10 {
		return strconv.Itoa(int(in))
	}
	if in < base {
		return string(rune(87 + in))
	}

	numPlaces := 0

	for ; in > 0; in /= base {
		place := in % base
		if place < 10 {
			reversedResult.WriteString(strconv.Itoa(int(place)))
		} else {
			reversedResult.WriteString(sr(place))
		}
		numPlaces++
	}

	if len(reversedResult.String()) < numPlaces {
		reversedResult.WriteString("1")
	}

	var result strings.Builder

	for i := len(reversedResult.String()) - 1; i >= 0; i-- {
		result.WriteString(string(reversedResult.String()[i]))
	}

	return result.String()
}

// sr takes an int value that should be between 10 and 35
// I'm not about to put checks in for that, just use it properly
func sr(val int64) string {
	return string(rune(runeDifference + val))
}

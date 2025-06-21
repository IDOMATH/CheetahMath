package base_change

import "github.com/IDOMATH/CheetahMath/formulas"

const runeDifference = 87

func ToTen(base int, bts string) int64 {
	var baseTen int64 = 0
	for i := 0; i < len(bts); i++ {
		char := rune(bts[len(bts)-1-i])
		if char-'0' > 9 {
			char = char - runeDifference
		} else {
			char = char - '0'
		}
		baseTen = baseTen + (int64(char) * formulas.IntPow(int64(base), int64(i)))
	}
	return baseTen
}

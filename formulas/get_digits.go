package formulas

import "strconv"

func GetDigits(num int) int {
	if num < 0 {
		num *= -1
	}
	return len(strconv.Itoa(num))
}

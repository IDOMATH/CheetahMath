package formulas

func GetDigits(num int) int {
	if num < 0 {
		num *= -1
	}
	digits := 0
	for ok := true; ok; ok = (num > 0) {
		num /= 10
		digits++
	}
	return digits
}

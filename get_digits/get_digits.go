package get_digits

func GetDigits(num int64) int64 {
	if num < 0 {
		num *= -1
	}
	var digits int64 = 0
	for ok := true; ok; ok = (num > 0) {
		num /= 10
		digits++
	}
	return digits
}

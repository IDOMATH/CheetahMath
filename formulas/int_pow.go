package formulas

// IntPow does the math.Pow() thing but only accepts integers and return an integer
// 0 is returned instead of an error.  Treat 0 as an error.
func IntPow(base, exponent int64) int64 {
	if exponent < 0 {
		return 0
	}
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}

	val := base
	var i int64
	for i = 2; i <= exponent; i++ {
		val *= base
	}
	return val
}

package formulas

import "errors"

func IntPow(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("only positive exponents are allowed")
	}
	if exponent == 0 {
		return 1, nil
	}
	if exponent == 1 {
		return base, nil
	}

	val := base
	for i := 2; i <= exponent; i++ {
		val *= base
	}
	return val, nil
}

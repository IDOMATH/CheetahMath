package formulas

import "testing"

func TestIntPow(t *testing.T) {
	var expected int64 = 1
	got := IntPow(2, 0)

	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}

	expected = 2
	got = IntPow(2, 1)
	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}

	expected = 8
	got = IntPow(2, 3)
	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}

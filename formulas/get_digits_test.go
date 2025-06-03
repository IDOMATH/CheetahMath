package formulas

import "testing"

func TestGetDigits(t *testing.T) {
	expected := 1
	got := GetDigits(3)
	if expected != got {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 1
	got = GetDigits(-5)
	if expected != got {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 3
	got = GetDigits(-345)
	if expected != got {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 5
	got = GetDigits(10000)
	if expected != got {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 1
	got = GetDigits(0)
	if expected != got {
		t.Error("Expected : ", expected, " but got: ", got)
	}
}

package base_change

import "testing"

func Test_ToTen_From36(t *testing.T) {
	var expected int64 = 10
	got := ToTen(36, "a")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 0
	got = ToTen(36, "0")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 35
	got = ToTen(36, "z")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 8
	got = ToTen(36, "8")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 36
	got = ToTen(36, "10")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 1295
	got = ToTen(36, "zz")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 109
	got = ToTen(36, "31")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 1296
	got = ToTen(36, "100")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 1260
	got = ToTen(36, "z0")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}
}

func Test_ToTen_From2(t *testing.T) {
	var expected int64 = 2
	got := ToTen(2, "10")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 10
	got = ToTen(2, "1010")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}
}

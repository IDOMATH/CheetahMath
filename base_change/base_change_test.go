package base_change

import "testing"

var table = []struct {
	in       string
	expected int64
}{
	{"a", 10},
	{"A", 10},
	{"0", 0},
	{"z", 35},
	{"8", 8},
	{"10", 36},
	{"zz", 1295},
	{"ZZ", 1295},
	{"Zz", 1295},
	{"31", 109},
	{"100", 1296},
	{"z0", 1260},
}

func Test_ToTen_From36(t *testing.T) {
	for _, testCase := range table {
		expected := testCase.expected
		got := ToTen(36, testCase.in)
		if got != expected {
			t.Error("expected: ", expected, "but got: ", got)
		}

	}
}

var fromTwoTable = []struct {
	in       string
	expected int64
}{
	{"10", 2},
	{"1010", 10},
}

func Test_ToTen_From2(t *testing.T) {
	for _, testCase := range fromTwoTable {
		expected := testCase.expected
		got := ToTen(2, testCase.in)
		if got != expected {
			t.Error("expected: ", expected, "but got: ", got)
		}

	}
}

var toTen_FromThirtySixTable = []struct {
	in       string
	expected int64
}{
	{"0", 0},
	{"a", 10},
	{"z", 35},
	{"10", 36},
	{"5", 5},
	{"9", 9},
	{"2s", 100},
	{"100", 1296},
}

func Test_ToTen_FromThirtySix(t *testing.T) {
	for _, testCase := range toTen_FromThirtySixTable {
		expected := testCase.expected
		got := ToTen(36, testCase.in)
		if got != expected {
			t.Error("expected: ", expected, "but got: ", got)
		}
	}
}

var toTen_FromTwoTable = []struct {
	in       string
	expected int64
}{
	{"0", 0},
	{"1010", 10},
	{"10", 2},
}

func Test_ToTen_FromTwo(t *testing.T) {
	for _, testCase := range toTen_FromTwoTable {
		expected := testCase.expected
		got := ToTen(2, testCase.in)
		if got != expected {
			t.Error("expected: ", expected, "but got: ", got)
		}
	}
}

func TestSr(t *testing.T) {
	expected := "a"
	got := sr(10)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = "z"
	got = sr(35)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}
}

package formulas

import (
	"testing"
)

func TestHaversine(t *testing.T) {
	expected := 0.8750338183720852
	got := Haversine(51.5007, 0.1246, 40.6892, 74.0445)

	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}

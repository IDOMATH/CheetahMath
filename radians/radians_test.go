package radians

import (
	"math"
	"testing"
)

func TestRadians(t *testing.T) {
	expected := math.Pi / 2.0
	got := ToRadians(90.0)

	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}

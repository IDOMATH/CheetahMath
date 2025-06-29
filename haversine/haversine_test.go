package haversine

import (
	"testing"
)

func TestHaversine(t *testing.T) {
	// Pi/2 radians
	expected := 1.5707963267948963
	// Equator to north pole
	got := Haversine(0, 0, 90, 0)

	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}

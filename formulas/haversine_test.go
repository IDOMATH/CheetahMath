package formulas

import (
	"fmt"
	"testing"
)

func TestHaversine(t *testing.T) {
	// expected := true
	got := Haversine(51.5007, 0.1246, 40.6892, 74.0445)
	fmt.Println(got)
}

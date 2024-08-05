package main

import (
	"fmt"

	"github.com/IDOMATH/Convert/formulas"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(formulas.Haversine(0, 0, 90, 0))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 3.5
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("The area of a circle with radius %g is %g.\n", radius, area)
}

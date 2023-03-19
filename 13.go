package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Pow(float64(x), 2) + math.Pow(float64(y), 2)
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

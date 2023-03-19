package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = true
	MaxInt uint64     = 42
	z      complex128 = cmplx.Sqrt(-7 + 24i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

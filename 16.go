package main

import "fmt"

const (
	// Create a large number by shifting a 1 bit left 50 places.
	// In other words, the binary number that is 1 followed by 50 zeroes.
	Large = 1 << 50
	// Shift it right again 49 places, so we end up with 1<<1, or 2.
	Small = Large >> 49
)

func needInt(x int) int { return x*5 + 2 }
func needFloat(x float64) float64 {
	return x * 0.2
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Large))
}

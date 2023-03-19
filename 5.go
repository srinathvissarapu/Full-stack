package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	
	result1 := add(5, 7)
	result2 := add(11, 22)

	
	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}

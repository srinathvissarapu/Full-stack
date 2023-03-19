package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	
	result1 := add(10, 20)
	result2 := add(30, 40)

	
	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}

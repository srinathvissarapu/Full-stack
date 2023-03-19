package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// Call the swap function with different values
	a, b := swap("hii", "srinath")
	c, d := swap("alpha", "beta")

	// Print the results
	fmt.Println(a, b)
	fmt.Println(c, d)
}

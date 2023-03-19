package main

import (
	"fmt"
	"math/rand"
)

func main() {
	
	randomNumber := rand.Intn(100) + 1

	
	if isEven(randomNumber) {
		fmt.Printf("%d is even\n", randomNumber)
	} else {
		fmt.Printf("%d is odd\n", randomNumber)
	}
}

func isEven(num int) bool {
	return num%2 == 0
}






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
In this program, we use the same fmt and math/rand packages as before, but we generate a random number between 1 and 100 using rand.Intn(100) + 1 instead of rand.Intn(10). We then call a custom function isEven to check if the number is even or odd, and print a message accordingly. This is a different method than the original code, which simply printed the random number.







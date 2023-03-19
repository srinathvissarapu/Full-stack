package main

import "fmt"

const Pi = 3.14159

func main() {
	const World = "Universe"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = false
	fmt.Println("Go rules?", Truth)
}

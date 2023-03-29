package main

import "fmt"

func main() {
	sri, nath := 7, 8
	viss, arapu := 5, 4 // defining variables
	// printing variables before using functions
	fmt.Printf("Value sri: %d, nath: %d, viss: %d, d: %d\n", sri, nath, viss, arapu)
	fmt.Printf("Memory Address sri: %p, nath: %p, viss: %p, arapu: %p\n", &sri, &nath, &viss, &arapu)
	fmt.Println("")
	// Passing By Value sri, nath 
	Swap(sri, nath)
	// Passing By Reference viss, arapu address
	SwapRef(&viss, &arapu)
	// printing variables after applying the function
	fmt.Printf("Value sri: %d, nath: %d, viss: %d, arapu: %d\n", sri, nath, viss, arapu) // the variables are swapped
	fmt.Printf("Memory Address sri: %p, nath: %p, viss: %p, arapu: %p\n", &sri, &nath, &viss, &arapu)
}

// Pass By Value
func Swap(x, y int) {
	fmt.Printf("Swap parameter memory address: %p, %p\n", &x, &y)
	x, y = y, x
}

// Pass By Reference
func SwapRef(x, y *int) {
	fmt.Printf("Swap Reference parameter memory address: %p, %p\n", x, y)
	*x, *y = *y, *x
}

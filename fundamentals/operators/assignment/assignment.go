package main

import "fmt"

func main() {
	var a byte = 3
	fmt.Println(a)

	// assignment types
	b := 10 // type inference
	b++     // b += 1 or b = b + 1
	b -= 2  // b = b -2
	b /= 3  // b = b / 3
	b *= 2  // b = b * 2
	b %= 5  // b = b % 5
	fmt.Println(b)

	// multiple statements, inference on the same line
	x, y := 5, 6
	fmt.Println(x, y)

	// change value
	x, y = y, x
	fmt.Println(x, y)
}

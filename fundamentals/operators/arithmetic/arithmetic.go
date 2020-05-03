package main

import (
	"fmt"
	"math"
)

func main() {
	// Simple operations
	a := 20
	b := 3
	fmt.Println("sum => ", a+b)
	fmt.Println("subtract => ", a-b)
	fmt.Println("divide => ", a/b)
	fmt.Println("multiply => ", a*b)
	fmt.Println("module => ", a%b)

	// Bitwise
	fmt.Println("E => ", a&b)
	fmt.Println("OR => ", a|b)
	fmt.Println("Xor => ", a^b)

	// Other operations using math...
	c := 30
	d := 7.5
	fmt.Println("highest value => ", math.Max(float64(c), d))
	fmt.Println("Lower value => ", math.Min(float64(c), d))
	fmt.Println("Exponential value => ", math.Pow(float64(c), d))
}

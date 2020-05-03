package main

import "fmt"

func main() {
	/*
		unary operators
	*/
	x := 1
	y := 2

	// only post fixed
	x++ // x += 1 or x = x + 1
	fmt.Println(x)

	y-- // y -= 1 or y = y - 1
	fmt.Println(y)

	/*
		WARNING...
		Don't do that
		fmt.Println(x == y--)
	*/
}

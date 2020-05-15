package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // type (float64) inferred by the compiler

	// reduced way to create a variable
	area := PI * m.Pow(raio, 2)

	fmt.Println("The circumference area is ", area)

	// you can declare consts in block
	const (
		a = 1
		b = 2
	)
	// you can declare variables in block
	var (
		c = 1
		d = 2
	)
	fmt.Println(a, b, c, d)

	// you can declare multiple variables on one line
	var e, f bool = true, false
	fmt.Println(e, f)

	// you can declare and initialize multiple variables on one line
	g, h, i := 2, false, "LOL"
	fmt.Println(g, h, i)
}

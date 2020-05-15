package main

import "fmt"

func main() {
	/*
		when we declare a variable of a certain type and do not assign a value,
		by default, an initial value is inferred in the declaration
	*/
	var a int     //infers the value 0 by default
	var b float64 // infers the value 0 by default
	var c bool    // infers the value false by default
	var d string  // infers an empty string
	var e *int    // infers the value <nil> by default

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}

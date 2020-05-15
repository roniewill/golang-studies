package main

import "fmt"

func main() {
	/*
		WARNING!!!
		No arithmetic pointer
	*/

	i := 1
	ma := &i
	ma2 := &ma
	ma3 := &i

	fmt.Println("Your value:", i)
	fmt.Println("first memory address:", ma)
	fmt.Println("second memory address:", ma2)
	fmt.Println("third memory address:", ma3)

	a := "memory"
	fmt.Println("Value of a:", a)

	var b *string = nil // nil == null
	fmt.Println("memory address of b before:", b)

	b = &a
	fmt.Println("memory address of b after:", b)

	// using deference, we can get the value of a variable
	fmt.Println("get value of memory address of b:", *b)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// conversion int to float64
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	// conversion float64 to int
	score := 6.9
	finalScore := int(score)
	fmt.Println(finalScore)

	// conversion int to string
	fmt.Println("Test string: " + strconv.Itoa(100))

	/*
		WARNING...
		This line will print value 'a' because 97 is a number that represents that character in the ASC table
	*/
	fmt.Println("Test: " + string(97))

	// conversion string to int
	num, _ := strconv.Atoi("200")
	fmt.Println(num - 50)

	// conversion string to bool
	boo, _ := strconv.ParseBool("true") // set "false" obviously prints false
	boo2, _ := strconv.ParseBool("1")   // set "0" it's equals false

	if boo {
		fmt.Println("is true")
	}

	if boo2 {
		fmt.Println("is true too")
	}
}

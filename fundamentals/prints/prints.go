package main

import "fmt"

func main() {
	// print on the same line
	fmt.Print("print on the")
	fmt.Print(" same line")
	fmt.Print("\n")

	// print on the new line
	fmt.Println("print on the")
	fmt.Println("new line")

	// print the formated value
	x := 3.1415
	fmt.Printf("The x value is: %f", x)
	fmt.Printf("\nThe x value is: %.2f", x)

	/*
		types print format
		%d to digit
		%f to float
		%t to bool
		%s to string
	*/
	a := 1
	b := 3.1415
	c := false
	d := "LOL"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}

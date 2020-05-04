package main

import "fmt"

func main() {
	fmt.Println("score average is: ", scoreAverage(9.5))
	fmt.Println("score average is: ", scoreAverage(8.3))
	fmt.Println("score average is: ", scoreAverage(7))
	fmt.Println("score average is: ", scoreAverage(6.1))
	fmt.Println("score average is: ", scoreAverage(4.5))
	fmt.Println("score average is: ", scoreAverage(1.8))
	fmt.Println("score average is: ", scoreAverage(33.5))
}

func scoreAverage(score float64) string {
	/*
		WARNING!
		switch clause does not automatically pass to the next case,
		on the contrary it is necessary to use fallthrough for this to happen.
	*/
	var average = int(score)
	switch average {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "invalid average"
	}
}

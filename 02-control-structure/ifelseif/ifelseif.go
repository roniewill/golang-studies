package main

import "fmt"

func main() {
	fmt.Println("score average is: ", scoreAverage(9.5))
	fmt.Println("score average is: ", scoreAverage(8.3))
	fmt.Println("score average is: ", scoreAverage(7))
	fmt.Println("score average is: ", scoreAverage(6.1))
	fmt.Println("score average is: ", scoreAverage(4.5))
}

func scoreAverage(score float64) string {
	if score >= 9 && score <= 10 {
		return "A"
	} else if score >= 8 && score < 9 {
		return "B"
	} else if score >= 7 && score < 8 {
		return "C"
	} else if score >= 5 && score < 7 {
		return "D"
	} else {
		return "E"
	}
}

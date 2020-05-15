package main

import (
	"fmt"
	"time"
)

func main() {
	// First Exampple
	/* 	fmt.Println("score average is: ", scoreAverage(9.5))
	   	fmt.Println("score average is: ", scoreAverage(8.3))
	   	fmt.Println("score average is: ", scoreAverage(7))
	   	fmt.Println("score average is: ", scoreAverage(6.1))
	   	fmt.Println("score average is: ", scoreAverage(4.5))
	   	fmt.Println("score average is: ", scoreAverage(1.8))
			 fmt.Println("score average is: ", scoreAverage(33.5)) */

	// Second Example
	// greetings()

	// Exercise
	/* 	fmt.Println("score average is: ", checkScoreAverage(9.5))
	   	fmt.Println("score average is: ", checkScoreAverage(8.3))
	   	fmt.Println("score average is: ", checkScoreAverage(7))
	   	fmt.Println("score average is: ", checkScoreAverage(6.1))
	   	fmt.Println("score average is: ", checkScoreAverage(4.5))
	   	fmt.Println("score average is: ", checkScoreAverage(1.8))
	   	fmt.Println("score average is: ", checkScoreAverage(33.5)) */

	// Third example
	fmt.Println(checkType(8.3))
	fmt.Println(checkType(func() {}))
	fmt.Println(checkType("Im string"))
	fmt.Println(checkType(100))
	fmt.Println(checkType(time.Now()))
}

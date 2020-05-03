package main

import "fmt"

/*
	WARNING!!! The Go programming language has no ternary operator.
	But, we will show an alternative
*/

func getResult(score float64) string {
	if score >= 6 {
		return "Approved"
	}
	return "disapproved"
}

func main() {
	score := 6.5
	fmt.Println(getResult(score))
}

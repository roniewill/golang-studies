package main

import "fmt"

func main() {
	showResult(6.9)
	showResult(7)
}

func showResult(score float64) {
	if score >= 7 {
		fmt.Println("approved")
	} else {
		fmt.Println("disapproved")
	}
}

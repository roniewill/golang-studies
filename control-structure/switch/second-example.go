package main

import (
	"fmt"
	"time"
)

func greetings() {
	t := time.Now() // get current date - obtem a data atual

	switch { // switch true {}
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternon!")
	default:
		fmt.Println("Good night!")
	}
}

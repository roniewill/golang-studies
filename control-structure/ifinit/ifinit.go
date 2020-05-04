package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	if i := randomNumbers(); i > 5 {
		fmt.Println("WINNER!")
	} else {
		fmt.Println("you lose...")
	}
}

func randomNumbers() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// first example
	i := 1
	for i <= 10 {
		fmt.Printf("%d", i)
		i++
	}

	// second example
	fmt.Println()
	for i := 0; i < 20; i += 2 {
		fmt.Printf("%d", i)
	}

	// third example
	fmt.Println()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Par %d; ", i)
		} else {
			fmt.Printf("Impar %d; ", i)
		}
	}

	// fourth example
	fmt.Println()
	for {
		fmt.Println("Infinit loop")
		time.Sleep(time.Second)
	}
}

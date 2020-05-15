package main

import "fmt"

func main() {
	/*
		PT_BR:
			Os arrays são estruturas homogêneas (mesmo tipo) e estática (fixo)

		ENG:
			Arrays are homogeneous (same type) and static (fixed) structures
	*/
	var score [3]float64
	fmt.Println(score)

	score[0], score[1], score[2] = 9.3, 7.5, 4.1
	fmt.Println(score)

	total := 0.0

	for i := 0; i < len(score); i++ {
		total += score[i]
	}
	average := total / float64(len(score))
	fmt.Printf("Average: %.2f", average)
}

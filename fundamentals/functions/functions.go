package main

import "fmt"

/*
	Go Lang by default builds/runs only the mentioned file.
	To Link all files you need to specify the name of all files while running.
	so you can run it: go run path/*.go
*/

func sum(a int, b int) int {
	return a + b
}

func print(value int) {
	fmt.Println(value)
}

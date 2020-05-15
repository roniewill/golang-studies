package main

/*
	Go Lang by default builds/runs only the mentioned file.
	To Link all files you need to specify the name of all files while running.
	so you can run it: go run path/*.go
*/

func main() {
	res := sum(10, 5)
	print(res)
}

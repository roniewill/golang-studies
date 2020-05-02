package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("studying the basic types")

	fmt.Println("whole literal is:", reflect.TypeOf(3000))

	/*
		whithout sinal only positives
		uint8  = 8 bits
		uint16 = 2 bytes
		uint32 = 4 bytes
		uint64 = 8 bytes
	*/
	var b byte = 255
	fmt.Println("The byte is:", reflect.TypeOf(b))

	/*
		whith sinal... int8  int16 int32 int64
	*/
	i1 := math.MaxInt64
	fmt.Println("The max value of int is:", i1)
	fmt.Println("The type of i1 is:", reflect.TypeOf(i1))

	var i2 = 'a' // represents a mapping in the Unicode table (int32)
	fmt.Println("The type of i2 is:", reflect.TypeOf(i2))
	fmt.Println("The value of i2 is:", i2)

	// real numbers (float32, float64)
	var x float32 = 49.99 // we can declare like that too, example: var x = float32(499.99)
	fmt.Println("The type of x is:", reflect.TypeOf(x))
	fmt.Println("The literal type is:", reflect.TypeOf(49.99)) // return a float64

	// boolean
	bo := true
	fmt.Println("The type bo is: ", reflect.TypeOf(bo))
	fmt.Println("The value bo is: ", bo)
	fmt.Println("The value bo is now: ", !bo)

	// string
	str := "Hello i'm a developer"
	fmt.Println(str + "!")
	fmt.Println("The length of str is:", len(str))

	str2 := `
	Hello!
	I'm
	Go
	developer
	beginner
	`
	fmt.Println(str2 + "!")
	fmt.Println("The length of str2 is:", len(str2))
}

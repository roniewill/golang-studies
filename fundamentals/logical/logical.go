package main

import "fmt"

func buy(job1, job2 bool) (bool, bool, bool) {
	buyTv50 := job1 && job2
	buyTv32 := job1 != job2
	buyIceCream := job1 || job2
	return buyTv50, buyTv32, buyIceCream
}

func main() {
	tv50, tv32, iceCream := buy(true, true)
	fmt.Printf("purchased Tv50: %t, purchased Tv32: %t, purchased Icecream: %t, you are healthy: %t", tv50, tv32, iceCream, !iceCream)
}

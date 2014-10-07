package main

import (
	"fmt"

	"./numbers"
)

func main() {

	for i := int64(220); i < 225; i++ {

		a := numbers.Divisors(i)
		a_sum := sum(a[:len(a)-1])
		b := numbers.Divisors(a_sum)
		b_sum := sum(b[:len(b)-1])

		fmt.Printf("i: %d a: %d b: %d\n", i, a_sum, b_sum)
	}
}

func sum(list []int64) int64 {

	sum := int64(0)

	for _, v := range list {
		sum += v
	}

	return sum
}

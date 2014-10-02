package main

import (
	"fmt"

	"./numbers"
)

func main() {

	max := 0

	for i := 1; ; i++ {

		sum := SumN(i)
		list := numbers.Divisors(int64(sum))

		if len(list) > max {
			fmt.Printf("Triangle %d has %d divisors.\n", sum, len(list))
			max = len(list)
		}

	}
}

func SumN(n int) int {
	return (n * (n + 1)) / 2
}

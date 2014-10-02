package main

import (
	"fmt"
)

func main() {

	max := 0

	for i := 1; ; i++ {

		sum := SumN(i)
		list := ListDivisors(sum)

		if len(list) > max {
			fmt.Printf("Triangle number %d has %d divisors.\n", sum, len(list))
			max = len(list)
		}

	}
}

func SumN(n int) int {
	return (n * (n + 1)) / 2
}

func ListDivisors(n int) []int {

	divisors := make([]int, 0)

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	divisors = append(divisors, n)
	return divisors
}

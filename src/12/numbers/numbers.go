package numbers

import "math"

func Divisors(n int64) []int64 {

	low := []int64{}
	high := []int64{}

	for i := int64(1); float64(i) < math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			low = append(low, i)
			high = append([]int64{n / i}, high...)
		}
	}

	return append(low, high...)
}

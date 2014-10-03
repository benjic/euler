package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])

	scanner := bufio.NewScanner(file)
	triangle := make([][]float64, 0)

	for scanner.Scan() {

		line := make([]float64, 0)

		for _, elem := range strings.Split(scanner.Text(), " ") {
			num, _ := strconv.ParseFloat(elem, 64)
			line = append(line, num)
		}

		triangle = append(triangle, line)
	}

	for row := len(triangle) - 2; row >= 0; row-- {
		for sub := 0; sub < len(triangle[row]); sub++ {
			triangle[row][sub] = math.Max(
				triangle[row][sub]+triangle[row+1][sub],
				triangle[row][sub]+triangle[row+1][sub+1])
		}
	}

	fmt.Printf("The maximized sum is %d\n", int(triangle[0][0]))
}

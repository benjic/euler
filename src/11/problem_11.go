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

	// Confirm input
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	data_file, err := os.Open(os.Args[1])

	// Confirm file given
	if err != nil {
		os.Exit(1)
	}

	// Build file and container
	scanner := bufio.NewScanner(data_file)
	grid := make([][]float64, 0)

	// Iterate over lines
	for scanner.Scan() {

		line := scanner.Text()
		row := make([]float64, 20)

		// Iterate over space seperated list of numbers
		for i, element := range strings.Split(line, " ") {

			number, err := strconv.ParseFloat(element, 64)

			// Confirm number parse
			if err != nil {
				fmt.Printf("Bad number conversion: %s", element)
				os.Exit(1)
			}

			row[i] = number
		}

		// Add row
		grid = append(grid, row)
	}

	// Start max at zero
	max := float64(0)

	// The product is computed using a 4x4 sliding index window
	// since multiplication is assoicative we need only try half
	// of directions.
	for row := 0; row < len(grid)-3; row++ {
		for col := 0; col < len(grid[row])-3; col++ {

			// Compute product of 4x1 horizontal and get max
			max = math.Max(grid[row][col]*grid[row][col+1]*grid[row][col+2]*grid[row][col+3], max)
			// compute produce of 1x4 vertical and get max
			max = math.Max(grid[row][col]*grid[row+1][col]*grid[row+2][col]*grid[row+3][col], max)
			// Compute left diaginal and get max
			max = math.Max(grid[row][col]*grid[row+1][col+1]*grid[row+2][col+2]*grid[row+3][col+3], max)
			// Compute right diaginal and get max
			max = math.Max(grid[row][col+3]*grid[row+1][col+2]*grid[row+2][col+1]*grid[row+3][col], max)

		}
	}

	// The end
	fmt.Printf("%d\n", int64(max))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])

	scanner := bufio.NewScanner(file)
	triangle := make([][]int64, 0)

	for scanner.Scan() {

		line := make([]int64, 0)

		for _, elem := range strings.Split(scanner.Text(), " ") {
			num, _ := strconv.ParseInt(elem, 10, 64)
			line = append(line, num)
		}

		triangle = append(triangle, line)
	}

	fmt.Printf("%v", triangle)
}

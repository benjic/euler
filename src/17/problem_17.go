package main

import (
	"fmt"
	"strings"
)

func main() {

	sum := 0

	for i := int64(1); i <= 1000; i++ {
		written_expression := BuildEnglishNumber(i)
		fmt.Printf("%s %d\n", written_expression, len(strings.Replace(written_expression, " ", "", -1)))
		sum += len(strings.Replace(written_expression, " ", "", -1))
	}

	fmt.Println(sum)
}

func BuildEnglishNumber(n int64) string {

	var ones, tens int64
	var number_str, split string

	if (n % 100) < 20 {
		tens = 0
		ones = n % 20
		split = ""
	} else {
		ones = n % 10
		tens = ((n % 100) - ones)
		split = ""
	}

	hundreds := ((n % 1000) - tens - (n % 10)) / 100
	thousands := (n - hundreds - tens - (n % 10)) / 1000

	if thousands > 0 {
		number_str += fmt.Sprintf("%sthousand ", NumericToWritten(thousands))
	}

	if hundreds > 0 {
		if tens > 0 || ones > 0 {
			number_str += fmt.Sprintf("%s hundred and ", NumericToWritten(hundreds))
		} else {
			number_str += fmt.Sprintf("%s hundred", NumericToWritten(hundreds))
		}
	}

	if tens > 0 {
		number_str += fmt.Sprintf("%s ", NumericToWritten(tens))
	}

	if ones > 0 {
		number_str += fmt.Sprintf("%s%s ", split, NumericToWritten(ones))
	}

	return strings.TrimRight(number_str, "- ")

}

func NumericToWritten(n int64) string {

	switch n {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	case 20:
		return "twenty "
	case 30:
		return "thirty "
	case 40:
		return "forty "
	case 50:
		return "fifty "
	case 60:
		return "sixty "
	case 70:
		return "seventy "
	case 80:
		return "eighty "
	case 90:
		return "ninety "

	}

	return ""
}

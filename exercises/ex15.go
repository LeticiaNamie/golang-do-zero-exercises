package exercises

import "strings"

func SortMyString(s string) string {
	arrayString := strings.SplitAfter(s, "")
	var evenHalf, oddHalf string

	for i := 0; i < len(arrayString); i++ {
		if EvenOrOdd2(i) == "Even" {
			evenHalf += arrayString[i]
		} else {
			oddHalf += arrayString[i]
		}
	}

	return evenHalf + " " + oddHalf
}

func EvenOrOdd2(number int) string {
	if number%2 == 0 {
		return "Even"
	}

	return "Odd"
}

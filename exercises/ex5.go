package exercises

import "strconv"

func StringToNumber(str string) int {
	numInt, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return numInt
}

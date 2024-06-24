package exercises

import "fmt"

func CalculateYears(years int) (result [3]int) {
	result[0] = years
	result[1] = 15
	result[2] = 15

	for i := 2; i <= years; i++ {
		if i == 2 {
			result[1] += 9
			result[2] = result[1]
		} else {
			result[1] += 4
			result[2] += 5
		}

		fmt.Println(result)
	}

	return
}

package exercises

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int

	res = append(res, 0, 0)

	for _, number := range numbers {
		if number > 0 {
			res[0] += 1
		} else {
			res[1] += number
		}
	}

	return res
}

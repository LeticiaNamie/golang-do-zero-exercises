package exercises

func Summation(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		total += i
	}

	return total
}

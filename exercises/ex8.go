package exercises

func Invert(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = -1 * arr[i]
	}

	return arr
}

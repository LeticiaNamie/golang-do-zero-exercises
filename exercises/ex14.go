package exercises

func PartList(arr []string) string {
	var answer, beforeComma, afterComma string
	var comma int = len(arr) - 1

	for i := 0; i < comma; i++ {
		beforeComma = ""
		afterComma = ""

		for j := 0; j <= i; j++ {
			if i == j {
				beforeComma += arr[j]
			} else {
				beforeComma += arr[j] + " "
			}
		}

		for k := i + 1; k < len(arr); k++ {
			afterComma += " "
			afterComma += arr[k]
		}

		answer += "(" + beforeComma + "," + afterComma + ")"
	}

	return answer
}

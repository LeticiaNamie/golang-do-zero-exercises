package exercises

func Well(x []string) string {
	var countGood, countBad int

	for _, idea := range x {
		if idea == "good" {
			countGood++
		} else {
			countBad++
		}

		if countGood > 2 {
			return "I smell a series!"
		}
	}

	if countGood == 1 || countGood == 2 {
		return "Publish!"
	}

	return "Fail!"
}

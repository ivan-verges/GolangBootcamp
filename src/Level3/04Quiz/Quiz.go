package quiz

func FilterAge(min int, max int, data []int) []int {

	result := []int{}

	for _, value := range data {
		if IsBetweenRange(value, min, max) {
			result = append(result, value)
		}
	}

	return result
}

func IsBetweenRange(number int, min int, max int) bool {
	if number >= min && number <= max {
		return true
	} else {
		return false
	}
}

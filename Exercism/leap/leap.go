package leap

func IsLeapYear(year int) bool {
	result := false

	if year % 400 == 0 {
		result = true
	} else if year % 100 == 0 {
		result = false
	} else if year % 4 == 0 {
		result = true
	}

	return result
}

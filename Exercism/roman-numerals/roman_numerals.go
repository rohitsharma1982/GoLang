package romannumerals

import "fmt"

// ToRomanNumeral converts a given arabic number to its Roman equivalent
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3000 {
		return "", fmt.Errorf("input outside the range")
	}

	result := ""
	changePointsInArabic := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	changePointsInRoman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for aitr, apoints := range changePointsInArabic {
		for arabic >= apoints {
			arabic -= apoints
			result += changePointsInRoman[aitr]
		}
	}

	return result, nil
}

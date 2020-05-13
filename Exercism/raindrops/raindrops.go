package raindrops

import "strconv"

// Convert converts a number into a string that contains raindrop sounds corresponding to certain potential factors.
// If the number has 3 as a factor, sound will include "Pling"
// If the number has 5 as a factor, sound will include "Plang"
// If the number has 7 as a factor, sound will include "Plong"
// If the number does not have any of 3, 5, or 7, the result will only have digits
func Convert(input int) string {
	var result string

	is3aFactor := input%3 == 0
	is5aFactor := input%5 == 0
	is7aFactor := input%7 == 0

	if !(is3aFactor || is5aFactor || is7aFactor) {
		return strconv.Itoa(input)
	}

	if is3aFactor {
		result += "Pling"
	}

	if is5aFactor {
		result += "Plang"
	}

	if is7aFactor {
		result += "Plong"
	}

	return result
}

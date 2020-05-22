package etl

import "strings"

// Transform takes scrabble scores from a legacy system and convert the scores to a new system.
//
// The old system stored a list of letters per score:
//
//- 1 point: "A", "E", "I", "O", "U", "L", "N", "R", "S", "T",
//- 2 points: "D", "G",
//- 3 points: "B", "C", "M", "P",
//- 4 points: "F", "H", "V", "W", "Y",
//- 5 points: "K",
//- 8 points: "J", "X",
//- 10 points: "Q", "Z",
//
// The shiny new Scrabble system instead stores the score per letter, which
// makes it much faster and easier to calculate the score for a word. It
// also stores the letters in lower-case regardless of the case of the
// input letters:
//
//- "a" is worth 1 point.
//- "b" is worth 3 points.
//- "c" is worth 3 points.
//- "d" is worth 2 points.
//- Etc.
func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)

	for key, value := range input {
		for _, item := range value {
			result[strings.ToLower(item)] = key
		}
	}
	return result
}

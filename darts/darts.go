package darts

import "math"

// Score returns the dart game score based on the distance of the dart from the center.
// Scoring is as follows:
//      1. 0 if outside the target (> 10)
//      2. 1 if in outer circle (> 5 && <= 10)
//      3. 5 if in middle circle (> 1 && <=5)
//      4. 10 if inside inner circle (< 1)
func Score(x, y float64) int {
	score := 0
	dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	if dist <= 1.0 {
		score = 10
	} else if dist <= 5.0 {
		score = 5
	} else if dist <= 10.0 {
		score = 1
	}

	return score
}

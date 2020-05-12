package hamming

import "fmt"

// Distance returns a hamming distance between two DNA strands of equal length.
// In case the length of strands is not equal, the function will throw an error.
func Distance(a, b string) (int, error) {

	var result int
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return 0, fmt.Errorf("unequal lengths: %q, %q", a, b)
	}

	for itr := range ar {
		if ar[itr] != br[itr] {
			result++
		}
	}

	return result, nil
}

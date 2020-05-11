package hamming

import "fmt"

// Distance returns a hamming distance between two DNA strands of equal length.
// In case the length of strands is not equal, the function will throw an error.
func Distance(a, b string) (int, error) {

	var result int = 0
	var err error = nil

	if len(a) != len(b) {
		err = fmt.Errorf("given dna strands are of different lengths")
		return result, err
	}

	for itr := range a {
		if a[itr] != b[itr] {
			result++
		}
	}

	return result, err
}

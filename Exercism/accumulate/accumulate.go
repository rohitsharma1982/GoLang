package accumulate

// Accumulate applies the input function on each element of the input collection and
// returns a collection with output of the function corresponding to each respective input element
func Accumulate(input []string, fn func(string) string) []string {
	result := make([]string, len(input))

	for itr, item := range input {
		result[itr] = fn(item)
	}

	return result
}

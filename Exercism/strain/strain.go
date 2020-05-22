package strain

// Ints represents a slice of ints
type Ints []int

// Lists represents a nested slice of ints i.e. a slice of slice of ints
type Lists [][]int

// Strings represents a slice of strings
type Strings []string

// Keep takes a predicate and returns a collection with only those elements of the input collection for which the predicate is true.
// Only for collection of integer lists.
func (input Ints) Keep(fn func(int) bool) Ints {
	var result Ints
	for itr := range input {
		if fn(input[itr]) {
			result = append(result, input[itr])
		}
	}
	return result
}

// Discard takes a predicate and returns a collection after discarding those elements for which the predicate is true
// Only for collection of integers
func (input Ints) Discard(fn func(int) bool) Ints {
	var result Ints
	for itr := range input {
		if !fn(input[itr]) {
			result = append(result, input[itr])
		}
	}
	return result
}

// Keep takes a predicate and returns a collection with only those elements of the input collection for which the predicate is true.
// Only for collection of integer lists
func (input Lists) Keep(fn func([]int) bool) Lists {
	var result Lists
	for itr := range input {
		if fn(input[itr]) {
			result = append(result, input[itr])
		}
	}
	return result
}

// Keep takes a predicate and returns a collection with only those elements of the input collection for which the predicate is true.
// Only for collection with strings
func (input Strings) Keep(fn func(string) bool) Strings {
	var result Strings
	for itr := range input {
		if fn(input[itr]) {
			result = append(result, input[itr])
		}
	}
	return result
}

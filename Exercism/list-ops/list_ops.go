package listops

// IntList type is equivalent to a Slice of Integers
type IntList []int

// unaryFunc is equivalent to a function that takes an integer argument and returns an integer
type unaryFunc func(int) int

// predFunc is equivalent of a predicate function that takes an integer argument and returns a boolean
type predFunc func(int) bool

// binFunc is equivalent of a accumulator function. To understand the details, please refer to the problem statement
type binFunc func(int, int) int

// Length counts the number of elements in a given list
func (list *IntList) Length() int { //Adding Length method to IntList
	listLength := 0

	for itr := range *list {
		listLength = itr + 1
	}

	return listLength
}

// Append takes a source list, adds all the members of the list passed as the argument,
// and returns the appended list containing elements from both the lists.
func (list *IntList) Append(appendList IntList) IntList { //Adding Append method to IntList
	firstListLength := list.Length()
	secondListLength := appendList.Length()

	newLength := firstListLength + secondListLength
	resultList := make(IntList, newLength)

	for itr, num := range *list {
		resultList[itr] = num
	}

	for itr, num := range appendList {
		resultList[itr+firstListLength] = num
	}

	return resultList
}

// Concat takes an array of lists are argument and returns a list with elements of all the lists appended to the source list.
func (list *IntList) Concat(lists []IntList) IntList { // Adding Concat method to IntList
	firstListLength := list.Length()

	newList := make(IntList, firstListLength)

	for itr, num := range *list {
		newList[itr] = num
	}

	for _, itrList := range lists {
		newList = newList.Append(itrList)
	}

	return newList
}

// Map takes a function as argument and returns a list with the function applied on each element of the input list.
func (list *IntList) Map(fn unaryFunc) IntList {
	firstListLength := list.Length()

	newList := make(IntList, firstListLength)

	for itr, num := range *list {
		newList[itr] = fn(num)
	}

	return newList
}

// Reverse returns an output list with the elements in reverse order of the input list.
func (list *IntList) Reverse() IntList {
	firstListLength := list.Length()

	newList := make(IntList, firstListLength)

	for itr, num := range *list {
		newList[firstListLength-itr-1] = num
	}

	return newList
}

// Filter takes a predicate function as input and returns a list with elements that passes (returns true) on the predicate function
func (list *IntList) Filter(fn predFunc) IntList {
	newList := make(IntList, 0)

	for _, num := range *list {
		if fn(num) {
			tempList := make(IntList, 1)
			tempList[0] = num
			newList = newList.Append(tempList)
		}
	}

	return newList
}

// Foldl , given a function, a list, and initial accumulator, folds each item into the accumulator from the left using function(accumulator, item)
func (list *IntList) Foldl(fn binFunc, init int) int {

	if list.Length() == 0 {
		return init
	}

	for _, num := range *list {
		init = fn(init, num)
	}

	return init
}

// Foldr , given a function, a list, and an initial accumulator, folds each item into the accumulator from the right using function(item, accumulator)
func (list *IntList) Foldr(fn binFunc, init int) int {
	listLength := list.Length()

	if listLength == 0 {
		return init
	}

	tempList := list.Reverse()

	for _, num := range tempList {
		init = fn(num, init)
	}

	return init
}

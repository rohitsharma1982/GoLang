package luhn

import (
	"strconv"
	"strings"
)

// Valid determines if the given input is valid or not as per Luhn's formula
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "") // Removing all spaces

	if len(input) <= 1 { // Length Check
		return false
	}

	s, err := strconv.Atoi(input) // In case of non-numerical data, the string to integer conversion will throw an error

	if err != nil {
		return false
	}

	var sum int
	var ss []int

	for s > 0 { // Storing the input string as a slice of ints
		rem := s % 10
		ss = append(ss, rem)
		s = (s - rem) / 10
	}

	if string(input[0]) == "0" { // String conversion to int removes the leading 0
		ss = append(ss, 0)
	}

	for itr := 1; itr < len(ss); itr += 2 { // Converting numbers higher than 9 as per the rules
		ss[itr] = ss[itr] * 2
		if ss[itr] > 9 {
			ss[itr] = ss[itr] - 9
		}
	}

	for itr := 0; itr < len(ss); itr++ {
		sum += ss[itr]
	}

	if sum%10 == 0 {
		return true
	}

	return false
}

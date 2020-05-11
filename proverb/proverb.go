// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	initialString := "For want of a %s the %s was lost."
	endString := "And all for the want of a %s."

	result := make([]string, 0)
	inputLength := len(rhyme)
	var appendString string

	if inputLength != 0 {
		for itr := 0; itr < inputLength-1; itr++ {
			tempString := fmt.Sprintf(initialString, rhyme[itr], rhyme[itr+1])
			result = append(result, tempString)
		}
		appendString = fmt.Sprintf(endString, rhyme[0])
		result = append(result, appendString)
	}
	return result
}

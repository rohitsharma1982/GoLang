// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	abbreviation := ""

	re := regexp.MustCompile(`[A-Za-z']*`)
	wordSlice := re.FindAllString(s, -1)

	if len(wordSlice) > 0 {
		for _, word := range wordSlice {
			if word != "" {
				abbreviation += strings.ToUpper(string(word[0]))
			}
		}
	}
	return abbreviation
}

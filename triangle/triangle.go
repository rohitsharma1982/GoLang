// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "Not a triangle."          // not a triangle
	Equ = "An equilateral triangle." // equilateral
	Iso = "An isoceles triangle."    // isosceles
	Sca = "An scalene triangle."     // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind // Not needed, but being used since it was given in the question

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) { // Checking if any of the side given is NaN
		k = NaT
		return k
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) { // Checking if any of the side given is +Inf or -Inf
		k = NaT
		return k
	}

	if (a <= 0) || (b <= 0) || (c <= 0) { // Checking if any of the side is less than 0
		k = NaT
		return k
	}

	if (a+b < c) || (b+c < a) || (a+c < b) { // Checking if some of any of the two sides is less than the third side
		k = NaT
		return k
	}

	if (a == b) && (b == c) { // Check for Equilateral triangle
		k = Equ
		return k
	} else if (a != b) && (b != c) && (c != a) { // Check for Scalene triangle
		k = Sca
		return k
	} else { // Else triangle is Isoceles
		k = Iso
		return k
	}
}

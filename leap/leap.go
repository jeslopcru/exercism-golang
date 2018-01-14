// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(aNumber int) bool {

	isDivisibleByFour := aNumber%4 == 0
	isNotDivisibleByOneHundred := aNumber%100 != 0
	isDivisibleByFourHundred := aNumber%400 == 0

	return isDivisibleByFour && (isNotDivisibleByOneHundred || isDivisibleByFourHundred)
}

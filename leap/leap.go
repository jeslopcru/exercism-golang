package leap

func IsLeapYear(aNumber int) bool {

	isDivisibleByFour := aNumber%4 == 0
	isNotDivisibleByOneHundred := aNumber%100 != 0
	isDivisibleByFourHundred := aNumber%400 == 0

	return isDivisibleByFour && (isNotDivisibleByOneHundred || isDivisibleByFourHundred)
}

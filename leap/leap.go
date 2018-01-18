package leap

func IsLeapYear(aNumber int) bool {

	notDivisibleFor100 := isDivisible(aNumber, 100) == false
	return isDivisible(aNumber, 4) && (notDivisibleFor100 || isDivisible(aNumber, 400))
}

func isDivisible(a int, b int) bool {
	return a%b == 0
}

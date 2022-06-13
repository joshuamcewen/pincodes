package pincodes

// splitDigits parses an integer and returns an array of integers representing each digit.
func splitDigits(number int) []int {
	var digits []int
	n := number

	if n == 0 {
		digits = []int{0}
	}

	for n != 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	return digits
}

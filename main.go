package pincodes

import (
	"sort"
)

// splitDigits parses an integer and returns an array of integers representing each digit.
func splitDigits(pin string) []string {
	var digits []string

	for _, digit := range pin {
		digits = append(digits, string(digit))
	}

	return digits
}

var keypad = [4][3]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
	{" ", "0", " "},
}

// getIndexes take a keypad digit and returns the row/column indexes.
func getIndexes(digit string) (int, int) {
	var rows, columns = len(keypad), len(keypad[0])
	var r, c = 0, 0

	for r < rows && c < columns {
		if keypad[r][c] == digit {
			break
		}

		if c == columns-1 {
			c = 0
			r++
		} else {
			c++
		}
	}

	return r, c
}

// getPossibleDigits takes a keypad digit and returns all possible selections.
func getPossibleDigits(digit string) []string {
	var digits = []string{digit}

	var rows, columns = len(keypad), len(keypad[0])
	var r, c = getIndexes(digit)

	// Left
	if c-1 >= 0 && keypad[r][c-1] != " " {
		left := keypad[r][c-1]
		digits = append(digits, left)
	}

	// Right
	if c+1 <= columns-1 && keypad[r][c+1] != " " {
		right := keypad[r][c+1]
		digits = append(digits, right)
	}

	// Up
	if r-1 >= 0 {
		up := keypad[r-1][c]
		digits = append(digits, up)
	}

	// Down
	if r+1 <= rows-1 && keypad[r+1][c] != " " {
		down := keypad[r+1][c]
		digits = append(digits, down)
	}

	return digits
}

var digitsByWear = []string{"5", "1", "2", "8", "7", "4", "0", "3", "6", "9"}

// getDigitPriority takes a number and generates a numeric priority based on wear.
func getDigitPriority(pin string) int {
	digits := splitDigits(pin)
	priority := 0

	for _, digit := range digits {
		for i, v := range digitsByWear {
			if v == digit {
				priority += len(digitsByWear) - i
			}
		}
	}

	return priority
}

// getPermutations takes a keypad combination and returns all possible permutations, most likely first.
func getPermutations(pin string) []string {
	digits := splitDigits(pin)
	var combinations [][]string

	for _, digit := range digits {
		combinations = append(combinations, getPossibleDigits(digit))
	}

	var accumulator = combinations[0]

	// For each array of digits...
	for i := 1; i < len(combinations); i++ {
		var updated []string

		//	For each digit in that array...
		for _, n := range combinations[i] {
			// For each number already in the accumulator...
			for _, a := range accumulator {
				updated = append(updated, a+n)
			}
		}

		accumulator = updated
	}

	sort.Slice(accumulator, func(i int, j int) bool {
		return getDigitPriority(accumulator[i]) > getDigitPriority(accumulator[j])
	})
	return accumulator
}

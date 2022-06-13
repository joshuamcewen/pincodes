package pincodes

import (
	"sort"
	"strconv"
)

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

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}

var keypad = [4][3]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
	{-1, 0, -1},
}

func getIndexes(number int) (int, int) {
	var rows, columns = len(keypad), len(keypad[0])
	var r, c = 0, 0

	for r < rows && c < columns {
		if keypad[r][c] == number {
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

func getPossibleDigits(number int) []int {
	var digits = []int{number}

	var rows, columns = len(keypad), len(keypad[0])
	var r, c = getIndexes(number)

	// Left
	if c-1 >= 0 && keypad[r][c-1] != -1 {
		left := keypad[r][c-1]
		digits = append(digits, left)
	}

	// Right
	if c+1 <= columns-1 && keypad[r][c+1] != -1 {
		right := keypad[r][c+1]
		digits = append(digits, right)
	}

	// Up
	if r-1 >= 0 {
		up := keypad[r-1][c]
		digits = append(digits, up)
	}

	// Down
	if r+1 <= rows-1 && keypad[r+1][c] != -1 {
		down := keypad[r+1][c]
		digits = append(digits, down)
	}

	return digits
}

var digitsByWear = []int{5, 1, 2, 8, 7, 4, 0, 3, 6, 9}

func getDigitPriority(number int) int {
	digits := splitDigits(number)
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

func getPermutations(number int) []int {
	digits := splitDigits(number)
	var combinations [][]int

	for _, digit := range digits {
		combinations = append(combinations, getPossibleDigits(digit))
	}

	var accumulator = combinations[0]

	// For each array of digits...
	for i := 1; i < len(combinations); i++ {
		var updated []int

		//	For each digit in that array...
		for _, n := range combinations[i] {
			// For each number already in the accumulator...
			for _, a := range accumulator {
				combination, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(n))
				updated = append(updated, combination)
			}
		}

		accumulator = updated
	}

	sort.Slice(accumulator, func(i int, j int) bool {
		return getDigitPriority(accumulator[i]) > getDigitPriority(accumulator[j])
	})
	return accumulator
}

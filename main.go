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

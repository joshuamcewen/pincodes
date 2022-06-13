package pincodes

import (
	"reflect"
	"testing"
)

func TestSplitDigits(t *testing.T) {
	var cases = []struct {
		Number   int
		Expected []int
	}{
		{Number: 10, Expected: []int{0, 1}},
		{Number: 0, Expected: []int{0}},
		{Number: 1234, Expected: []int{4, 3, 2, 1}},
	}

	for _, c := range cases {
		got := splitDigits(c.Number)
		expected := c.Expected

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}
	}
}

func TestGetIndexes(t *testing.T) {
	var cases = []struct {
		Number   int
		Expected []int
	}{
		{Number: 1, Expected: []int{0, 0}},
		{Number: 5, Expected: []int{1, 1}},
		{Number: 9, Expected: []int{2, 2}},
		{Number: 0, Expected: []int{3, 1}},
	}

	for _, c := range cases {
		row, column := getIndexes(c.Number)
		expected := c.Expected

		if row != expected[0] || column != expected[1] {
			t.Errorf("Expected: %v, Got: %v", expected, []int{row, column})
		}
	}
}

func TestGetPossibleDigits(t *testing.T) {
	var cases = []struct {
		Number   int
		Expected []int
	}{
		{Number: 1, Expected: []int{1, 2, 4}},
		{Number: 5, Expected: []int{5, 4, 6, 2, 8}},
		{Number: 9, Expected: []int{9, 8, 6}},
		{Number: 0, Expected: []int{0, 8}},
	}

	for _, c := range cases {
		got := getPossibleDigits(c.Number)
		expected := c.Expected

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}
	}
}

func TestGetDigitPriority(t *testing.T) {
	var cases = []struct {
		Number   int
		Expected int
	}{
		{Number: 5, Expected: 10},
		{Number: 9, Expected: 1},
		{Number: 7, Expected: 6},
	}

	for _, c := range cases {
		got := getDigitPriority(c.Number)
		expected := c.Expected

		if expected != got {
			t.Errorf("Expected: %d, Got: %d", expected, got)
		}
	}
}

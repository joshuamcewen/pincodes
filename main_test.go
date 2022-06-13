package pincodes

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestSplitDigits(t *testing.T) {
	var cases = []struct {
		Pin      string
		Expected []string
	}{
		{Pin: "10", Expected: []string{"1", "0"}},
		{Pin: "0", Expected: []string{"0"}},
		{Pin: "1234", Expected: []string{"1", "2", "3", "4"}},
	}

	for _, c := range cases {
		got := splitDigits(c.Pin)
		expected := c.Expected

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}
	}
}

func TestGetIndexes(t *testing.T) {
	var cases = []struct {
		Digit    string
		Expected []int
	}{
		{Digit: "1", Expected: []int{0, 0}},
		{Digit: "5", Expected: []int{1, 1}},
		{Digit: "9", Expected: []int{2, 2}},
		{Digit: "0", Expected: []int{3, 1}},
	}

	for _, c := range cases {
		row, column := getIndexes(c.Digit)
		expected := c.Expected

		if row != expected[0] || column != expected[1] {
			t.Errorf("Expected: %v, Got: %v", expected, []int{row, column})
		}
	}
}

func TestGetPossibleDigits(t *testing.T) {
	var cases = []struct {
		Digit    string
		Expected []string
	}{
		{Digit: "1", Expected: []string{"1", "2", "4"}},
		{Digit: "5", Expected: []string{"5", "4", "6", "2", "8"}},
		{Digit: "9", Expected: []string{"9", "8", "6"}},
		{Digit: "0", Expected: []string{"0", "8"}},
	}

	for _, c := range cases {
		got := getPossibleDigits(c.Digit)
		expected := c.Expected

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}
	}
}

func TestGetDigitPriority(t *testing.T) {
	var cases = []struct {
		Digit    string
		Expected int
	}{
		{Digit: "5", Expected: 10},
		{Digit: "9", Expected: 1},
		{Digit: "7", Expected: 6},
	}

	for _, c := range cases {
		got := getDigitPriority(c.Digit)
		expected := c.Expected

		if expected != got {
			t.Errorf("Expected: %d, Got: %d", expected, got)
		}
	}
}

func TestGetPermutations(t *testing.T) {
	var cases = []struct {
		Pin      string
		Expected []string
	}{
		{Pin: "46", Expected: []string{"55", "15", "75", "45", "53", "56", "13", "59", "16", "19", "73", "43", "76", "46", "79", "49"}},
		{Pin: "1", Expected: []string{"1", "2", "4"}},
		{Pin: "8", Expected: []string{"5", "8", "7", "0", "9"}},
		{Pin: "01", Expected: []string{"81", "82", "01", "02", "84", "04"}},
	}

	for _, c := range cases {
		got := getPermutations(c.Pin)
		expected := c.Expected

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}
	}
}

func TestGetPermutations97516(t *testing.T) {
	permutations := getPermutations("97516")
	json, _ := json.Marshal(permutations)

	fmt.Printf("%s", json)
}

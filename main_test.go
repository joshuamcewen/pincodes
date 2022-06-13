package pincodes

import (
	"reflect"
	"testing"
)

// this: is a func
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

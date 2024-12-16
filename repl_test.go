package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Length of returned slice (%d) does not match expected length (%d)", len(actual), len(expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message and fail the test
			if word != expectedWord {
				t.Errorf("cleanInput(%s) == %s, expected %s", c.input, actual, expected)
			}
		}
	}
}

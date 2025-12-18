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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
        // This is the new test case for your explore command
		{
			input:    "EXPLORE canalave-city-area",
			expected: []string{"explore", "canalave-city-area"},
		},
        // Test ignoring empty input
        {
            input:    "",
            expected: []string{},
        },
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the slice
		if len(actual) != len(c.expected) {
			t.Errorf("len(actual) == %v, expected %v", len(actual), len(c.expected))
            continue
		}
        // Check each word
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "HEyhey hey    ",
			expected: []string{"heyhey", "hey"},
		},
		{
			input:    "     1A 223     b",
			expected: []string{"1a", "223", "b"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test fail! Length is not the same as expected. Values: %v, %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test fail! New word is not the same as expected")
			}
		}
	}

}

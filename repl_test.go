package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  tIger BEETLE World  ",
			expected: []string{"tiger", "beetle", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("got %q, expected %q", actual, c.expected)
			}
		}
	}
}

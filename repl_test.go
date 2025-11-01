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
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "My    name    is    Inigo Montoya     ",
			expected: []string{"my", "name", "is", "inigo", "montoya"},
		},
		{
			input:    "HeLlO wOrLd UwU",
			expected: []string{"hello", "world", "uwu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length %d(%v), got %d(%v)", len(c.expected), c.expected, len(actual), actual)
			t.FailNow()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word %s, got %s", expectedWord, word)
				t.FailNow()
			}
		}
	}
}

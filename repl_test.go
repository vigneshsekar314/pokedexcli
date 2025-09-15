package main

import "testing"

type CleanInputCase struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []CleanInputCase{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		}, {
			input:    "valid case test",
			expected: []string{"valid", "case", "test"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("words length incorrect. Expected: %d. Actual: %d.", len(cs.expected), len(actual))
		}
		for i, word := range cs.expected {
			if word != actual[i] {
				t.Errorf("words do not match. Expected - %s. Actual - %s", cs.expected[i], word)
			}

		}
	}

}

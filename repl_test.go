package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "this is not broken!",
			expected: []string{
				"this",
				"is",
				"not",
				"broken!",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal - %v vs %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("The values are not equal - %s vs %s", actualWord, expectedWord)
				continue
			}
		}
	}
}

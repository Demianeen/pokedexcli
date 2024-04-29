package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HEllo World",
			expected: []string{"hello", "world"},
		},
	}

	for _, cs := range cases {
		actual := getCleanedArgs(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Error with input %v", cs.input)
			t.Errorf("The length are not equal: expected length is %v, but got %v", cs.expected, actual)
			continue
		}
		for i := 0; i < len(actual); i++ {
			expectedArg := cs.expected[i]
			actualArg := actual[i]
			if actualArg != expectedArg {
				t.Errorf("Expected %v, but got %v", expectedArg, actualArg)
			}
		}
	}
}

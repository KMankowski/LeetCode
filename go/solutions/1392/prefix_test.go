package main

import (
	"fmt"
	"testing"
)

func TestLongestPrefix(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expLongestPrefix string
	}{
		{
			"example 1",
			"level",
			"l",
		},
		{
			"example 2",
			"ababab",
			"abab",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("naive: %s", test.name), func(t *testing.T) {
			outLongestPrefix := longestPrefixNaive(test.input)
			if outLongestPrefix != test.expLongestPrefix {
				t.Fatalf("got %v want %v", outLongestPrefix, test.expLongestPrefix)
			}
		})
	}
}

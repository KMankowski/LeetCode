package main

import (
	"slices"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name            string
		inpNums         []int
		inpK            int
		expMostFrequent []int
	}{
		{
			"example 1",
			[]int{1, 2, 2, 3, 3, 3},
			2,
			[]int{2, 3},
		},
		{
			"example 2",
			[]int{7, 7},
			1,
			[]int{7},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outMostFrequent := topKFrequent(test.inpNums, test.inpK)

			// Ordering of solution does not matter
			sort.Ints(outMostFrequent)
			sort.Ints(test.expMostFrequent)

			if !slices.Equal(outMostFrequent, test.expMostFrequent) {
				t.Fatalf("got %v want %v", outMostFrequent, test.expMostFrequent)
			}
		})
	}
}

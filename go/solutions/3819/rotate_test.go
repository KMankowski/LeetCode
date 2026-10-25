package main

import (
	"slices"
	"testing"
)

func TestRotateElements(t *testing.T) {
	tests := []struct {
		name    string
		inpNums []int
		inpK    int
		expNums []int
	}{
		{
			"example 1",
			[]int{1, -2, 3, -4},
			3,
			[]int{3, -2, 1, -4},
		},
		{
			"example 2",
			[]int{-3, -2, 7},
			1,
			[]int{-3, -2, 7},
		},
		{
			"example 3",
			[]int{5, 4, -9, 6},
			2,
			[]int{6, 5, -9, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outNums := rotateElements(test.inpNums, test.inpK)

			if !slices.Equal(outNums, test.expNums) {
				t.Fatalf("got %v want %v", outNums, test.expNums)
			}
		})
	}
}

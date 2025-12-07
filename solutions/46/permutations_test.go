package main

import (
	"slices"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name            string
		inpNums         []int
		expPermutations [][]int
	}{
		{
			"example 1",
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			"example 2",
			[]int{0, 1},
			[][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			"example 3",
			[]int{1},
			[][]int{
				{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			permutations := permute(test.inpNums)

			if len(permutations) != len(test.expPermutations) {
				t.Fatalf("expected %v got %v", test.expPermutations, permutations)
			}

			for _, expPermutation := range test.expPermutations {
				isInPermutations := false
				for _, permutation := range permutations {
					if slices.Equal(permutation, expPermutation) {
						isInPermutations = true
					}
				}
				if !isInPermutations {
					t.Fatalf("expected %v got %v", test.expPermutations, permutations)
				}
			}
		})
	}
}

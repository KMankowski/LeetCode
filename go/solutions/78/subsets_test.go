package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name       string
		nums       []int
		expSubsets [][]int
	}{
		{
			"example",
			[]int{1, 2, 3},
			[][]int{
				{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3},
			},
		},
		{
			"failing test",
			[]int{3, 2, 1, 4},
			[][]int{
				{}, {3}, {2}, {1}, {4}, {3, 2}, {3, 1}, {3, 4}, {2, 1}, {2, 4}, {1, 4},
				{3, 2, 1}, {3, 2, 4}, {3, 1, 4}, {2, 1, 4}, {3, 2, 1, 4},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			subsets := subsets(test.nums)

			sort := func(subsets [][]int) [][]int {
				for _, subset := range subsets {
					sort.Ints(subset)
				}
				sort.Slice(subsets, func(firstSubset, secondSubset int) bool {
					if len(subsets[firstSubset]) < len(subsets[secondSubset]) {
						return true
					} else if len(subsets[firstSubset]) > len(subsets[secondSubset]) {
						return false
					}
					for subsetIdx := range len(subsets[firstSubset]) {
						if subsets[firstSubset][subsetIdx] < subsets[secondSubset][subsetIdx] {
							return true
						} else if subsets[firstSubset][subsetIdx] > subsets[secondSubset][subsetIdx] {
							return false
						}
					}
					// doesn't matter at this point, since both []int will be such that
					// each slice will be of the same length and have the same values
					return true
				})
				return subsets
			}

			if !reflect.DeepEqual(sort(test.expSubsets), sort(subsets)) {
				t.Fatalf("expected %v got %v", test.expSubsets, subsets)
			}
		})
	}
}

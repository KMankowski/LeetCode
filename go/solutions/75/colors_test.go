package main

import (
	"slices"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name      string
		inpColors []int
		expColors []int
	}{
		{
			"example 1",
			[]int{1, 0, 1, 2},
			[]int{0, 1, 1, 2},
		},
		{
			"example 2",
			[]int{2, 1, 0},
			[]int{0, 1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortColors(test.inpColors)

			if !slices.Equal(test.inpColors, test.expColors) {
				t.Fatalf("got %v want %v", test.inpColors, test.expColors)
			}
		})
	}
}

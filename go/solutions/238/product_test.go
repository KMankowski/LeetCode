package main

import (
	"slices"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name        string
		inpNums     []int
		expProducts []int
	}{
		{
			"example 1",
			[]int{1, 2, 4, 6},
			[]int{48, 24, 12, 8},
		},
		{
			"example 2",
			[]int{-1, 0, 1, 2, 3},
			[]int{0, -6, 0, 0, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outProducts := productExceptSelf(test.inpNums)
			if !slices.Equal(outProducts, test.expProducts) {
				t.Fatalf("got %v want %v", outProducts, test.expProducts)
			}
		})
	}
}

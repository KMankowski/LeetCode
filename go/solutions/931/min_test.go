package main

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	tests := []struct {
		name   string
		m      [][]int
		expSum int
	}{
		{
			"example 1",
			[][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
			13,
		},
		{
			"nil input",
			nil,
			0,
		},
		{
			"nil row",
			[][]int{{2, 1, 3}, nil, {7, 8, 9}},
			0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outSum := minFallingPathSum(test.m)
			if outSum != test.expSum {
				t.Fatalf("got %v want %v", outSum, test.expSum)
			}
		})
	}
}

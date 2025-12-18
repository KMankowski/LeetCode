package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name         string
		inpIntervals [][]int
		expIntervals [][]int
	}{
		{
			"example 1",
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"example 2",
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			"example 3",
			[][]int{{4, 7}, {1, 4}},
			[][]int{{1, 7}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outIntervals := merge(test.inpIntervals)
			if !reflect.DeepEqual(outIntervals, test.expIntervals) {
				t.Fatalf("got %v want %v", outIntervals, test.expIntervals)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name         string
		intervals    [][]int
		newInterval  []int
		expIntervals [][]int
	}{
		{
			"no intersection, prepend",
			[][]int{{2, 3}, {6, 7}},
			[]int{0, 1},
			[][]int{{0, 1}, {2, 3}, {6, 7}},
		},
		{
			"no intersection, postpend",
			[][]int{{2, 3}, {6, 7}},
			[]int{8, 9},
			[][]int{{2, 3}, {6, 7}, {8, 9}},
		},
		{
			"no intersection, between two existing intervals",
			[][]int{{2, 3}, {6, 7}},
			[]int{4, 5},
			[][]int{{2, 3}, {4, 5}, {6, 7}},
		},
		{
			"intersection, new[end] == existing[start]",
			[][]int{{2, 3}, {6, 7}},
			[]int{1, 2},
			[][]int{{1, 3}, {6, 7}},
		},
		{
			"intersection, new[end] > existing[end]",
			[][]int{{2, 4}, {6, 7}},
			[]int{1, 3},
			[][]int{{1, 4}, {6, 7}},
		},
		{
			"intersection, new[end] == existing[end]",
			[][]int{{2, 4}, {6, 7}},
			[]int{1, 4},
			[][]int{{1, 4}, {6, 7}},
		},
		{
			"intersection, new is wrapped in existing",
			[][]int{{2, 4}, {6, 7}},
			[]int{3, 4},
			[][]int{{2, 4}, {6, 7}},
		},
		{
			"intersection, existing is wrapped in new",
			[][]int{{2, 4}, {6, 7}},
			[]int{1, 5},
			[][]int{{1, 5}, {6, 7}},
		},
		{
			"intersection, several merges",
			[][]int{{2, 4}, {6, 7}},
			[]int{1, 6},
			[][]int{{1, 7}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outIntervals := insert(test.intervals, test.newInterval)
			if !reflect.DeepEqual(outIntervals, test.expIntervals) {
				t.Fatalf("expected %v got %v", test.expIntervals, outIntervals)
			}
		})
	}
}

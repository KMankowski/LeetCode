package main

import (
	"fmt"
	"sort"
)

const (
	start = 0
	end   = 1
)

func main() {
	fmt.Println(merge([][]int{{4, 7}, {1, 4}}))
}

func merge(intervals [][]int) [][]int {
	// sort by ascending start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][start] < intervals[j][start]
	})

	var newIntervals [][]int
	i := 0
	for {
		if i == len(intervals) {
			break
		}

		j := i + 1
		for j < len(intervals) && intervals[i][end] >= intervals[j][start] {
			intervals[i][end] = max(intervals[i][end], intervals[j][end])
			j++
		}
		newIntervals = append(newIntervals, intervals[i])
		i = j
	}
	return newIntervals
}

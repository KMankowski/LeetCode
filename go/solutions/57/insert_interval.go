package main

import (
	"fmt"
)

const (
	start = 0
	end   = 1
)

func main() {
	fmt.Println(insert([][]int{[]int{1, 3}, []int{6, 9}}, []int{2, 5}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	newIntervals := make([][]int, len(intervals)+1)

	for i := range intervals {
		// loop while the newInterval is entirely offset to the right of each current interval
		if newInterval[start] > intervals[i][end] {
			newIntervals[i] = intervals[i]
			continue
		}

		// newInterval fits cleanly either at the head, or in between two intervals
		if isDisjoint(newInterval, intervals[i]) {
			newIntervals[i] = newInterval
			copy(newIntervals[i+1:], intervals[i:])
			return newIntervals
		}

		// now, we are intersecting an existing interval, so merge towards the right!
		newStart := min(newInterval[start], intervals[i][start])
		newEnd := max(newInterval[end], intervals[i][end])
		j := i + 1
		for j < len(intervals) && !isDisjoint([]int{newStart, newEnd}, intervals[j]) {
			newStart = min(newStart, intervals[j][start])
			newEnd = max(newEnd, intervals[j][end])
			j++
		}
		newIntervals[i] = []int{newStart, newEnd}
		copy(newIntervals[i+1:], intervals[j:])
		newIntervals = newIntervals[:i+1+len(intervals)-j]
		return newIntervals
	}

	// the start of the newInterval was beyond the end of each existing interval, postpend
	newIntervals[len(newIntervals)-1] = newInterval
	return newIntervals
}

func isDisjoint(i, j []int) bool {
	return i[end] < j[start] || i[start] > j[end]
}

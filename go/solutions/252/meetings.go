package main

import (
	"fmt"
	"sort"
)

type interval struct {
	start int
	end   int
}

func main() {
	fmt.Println(canAttendMeetings([]interval{{0, 30}, {5, 10}, {15, 20}}))
}

func canAttendMeetings(intervals []interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < intervals[i-1].end {
			return false
		}
	}
	return true
}

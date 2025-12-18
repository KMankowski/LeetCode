package main

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals []interval
		exp       bool
	}{
		{
			"example 1",
			[]interval{{15, 20}, {5, 10}, {0, 30}},
			false,
		},
		{
			"example 2",
			[]interval{{5, 8}, {9, 15}},
			true,
		},
		{
			"edge",
			[]interval{{5, 8}, {8, 15}},
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := canAttendMeetings(test.intervals)
			if test.exp != out {
				t.Fatalf("expected %v got %v", test.exp, out)
			}
		})
	}
}

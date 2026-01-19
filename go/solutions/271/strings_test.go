package main

import (
	"slices"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name    string
		inpStrs []string
	}{
		{
			"example 1",
			[]string{"Hello", "World"},
		},
		{
			"example 2",
			[]string{""},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := new(Solution)
			outStrs := s.Decode(s.Encode(test.inpStrs))
			if !slices.Equal(test.inpStrs, outStrs) {
				t.Fatalf("want %v got %v", test.inpStrs, outStrs)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name       string
		inpStrs    []string
		expEncoded string
	}{
		{
			"example 1",
			[]string{"Hello", "World"},
			"5.Hello5.World",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := new(Solution)
			outEncoded := s.Encode(test.inpStrs)
			if outEncoded != test.expEncoded {
				t.Fatalf("got %v want %v", outEncoded, test.expEncoded)
			}
		})
	}
}

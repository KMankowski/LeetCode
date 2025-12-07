package main

import (
	"log/slog"
	"os"
)

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	nums := []int{1, 2, 3}
	permutations := permute(nums)

	l.Info("success", "permutations", permutations)
}

func permute(nums []int) [][]int {
	allPaths := make([][]int, 0)
	currPath := make([]int, 0)
	inuseIndices := make([]bool, len(nums))
	return recurse(allPaths, currPath, nums, inuseIndices)
}

func recurse(allPaths [][]int, currPath []int, nums []int, inuseIndices []bool) [][]int {
	if len(currPath) == len(nums) {
		permutation := make([]int, len(currPath))
		copy(permutation, currPath)
		return append(allPaths, permutation)
	}
	for i := range nums {
		if inuseIndices[i] {
			continue
		}
		inuseIndices[i] = true
		currPath = append(currPath, nums[i])

		allPaths = recurse(allPaths, currPath, nums, inuseIndices)

		currPath = currPath[:len(currPath)-1]
		inuseIndices[i] = false
	}
	return allPaths
}

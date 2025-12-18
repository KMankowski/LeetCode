package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	return recurse(nums, 0, []int{}, [][]int{})
}

func recurse(nums []int, nextIdx int, currSubset []int, subsets [][]int) [][]int {
	if nextIdx == len(nums) {
		newSubset := make([]int, len(currSubset))
		copy(newSubset, currSubset)
		return append(subsets, newSubset)
	}

	// do include current num
	currSubset = append(currSubset, nums[nextIdx])
	subsets = recurse(nums, nextIdx+1, currSubset, subsets)

	// don't include current num
	currSubset = currSubset[:len(currSubset)-1]
	subsets = recurse(nums, nextIdx+1, currSubset, subsets)

	return subsets
}

package main

func rotateElements(nums []int, k int) []int {
	newNums := make([]int, len(nums))

	nonNegativeNums := make([]int, 0)
	originalIndices := make([]int, 0)
	for i, num := range nums {
		if num < 0 {
			newNums[i] = num
		} else {
			nonNegativeNums = append(nonNegativeNums, num)
			originalIndices = append(originalIndices, i)
		}
	}

	if len(nonNegativeNums) == 0 {
		return newNums
	}

	k = k % len(nonNegativeNums)
	for i, num := range nonNegativeNums {
		if i-k >= 0 {
			newNums[originalIndices[i-k]] = num
		} else {
			newNums[originalIndices[i-k+len(nonNegativeNums)]] = num
		}
	}

	return newNums
}

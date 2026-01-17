package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 2, 2, 3, 3, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	numToFrequency := make(map[int]int)
	for _, num := range nums {
		numToFrequency[num]++
	}

	// size of len(nums) since that is the largest possible frequency
	// +1 since 0th index is ignored
	frequencyToNum := make([][]int, len(nums)+1)
	for num, frequency := range numToFrequency {
		frequencyToNum[frequency] = append(frequencyToNum[frequency], num)
	}

	var kMostFrequent []int
	for frequency := len(frequencyToNum) - 1; frequency >= 0; frequency-- {
		for _, num := range frequencyToNum[frequency] {
			if len(kMostFrequent) == k {
				return kMostFrequent
			}
			kMostFrequent = append(kMostFrequent, num)
		}
	}
	return kMostFrequent
}

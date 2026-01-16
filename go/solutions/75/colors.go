package main

const NUMBER_OF_COLORS = 3

func sortColors(nums []int) {
	buckets := make([]int, NUMBER_OF_COLORS)

	for _, num := range nums {
		buckets[num]++
	}

	insertionIndex := 0
	for color := range buckets {
		for range buckets[color] {
			nums[insertionIndex] = color
			insertionIndex++
		}
	}
}

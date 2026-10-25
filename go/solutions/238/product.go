package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 4, 6}))
}

func productExceptSelf(nums []int) []int {
	prefixProducts := make([]int, len(nums))
	postfixProducts := make([]int, len(nums))
	products := make([]int, len(nums))

	prefixProduct := 1
	for i := 0; i < len(nums); i++ {
		prefixProducts[i] = prefixProduct
		prefixProduct *= nums[i]
	}

	postfixProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		postfixProducts[i] = postfixProduct
		postfixProduct *= nums[i]
	}

	for i := range products {
		products[i] = prefixProducts[i] * postfixProducts[i]
	}

	return products
}

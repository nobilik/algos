package arrays

import "math"

// Given an array that contains both positive and negative integers, the task is to find the product of the maximum product subarray.
// Time Complexity: O(N)
// Auxiliary Space: O(1)
func MaxSubarrayProduct(arr []int) int {
	ans := math.MinInt64 // Initialize the answer to the minimum possible value | is it ok?
	product := 1

	for _, el := range arr {
		product *= el
		ans = max(ans, product) // Update the answer with the maximum of the current answer and product
		if el == 0 {
			product = 1 // Reset the product to 1 if the current element is 0
		}
	}

	product = 1

	for i := len(arr) - 1; i >= 0; i-- {
		product *= arr[i]
		ans = max(ans, product)
		if arr[i] == 0 {
			product = 1
		}
	}

	return ans
}

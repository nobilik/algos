package arrays

import (
	"algos/helpers"
)

// use pivot index
// Time Complexity: O(N)
// Auxiliary Space: O(1)
func EquilibriumIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, pivot, right := 0, 0, helpers.Sum(nums)-nums[0]
	for pivot < len(nums)-1 && right != left {
		pivot++
		right -= nums[pivot]
		left += nums[pivot-1]
	}

	if left == right {
		return pivot
	}
	return -1
}

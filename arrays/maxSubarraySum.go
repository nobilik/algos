package arrays

// "algos/helpers"

// Given an array arr[] of size N. The task is to find the sum of the contiguous subarray within a arr[] with the largest sum.
// Kadane’s algorithm
// Time Complexity: O(n)
// Auxiliary Space: O(1)

// sum only
func MaxSubarraySumSimple(arr []int) int {
	// in case of empty array
	if len(arr) == 0 {
		return 0
	}

	maxSoFar := arr[0]
	maxEndingHere := arr[0]

	for _, num := range arr[1:] {
		maxEndingHere = max(maxEndingHere+num, num)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

// with indexes
func MaxSubarraySum(nums []int) (int, int, int) {
	// in case of empty array
	if len(nums) == 0 {
		return 0, -1, -1
	}

	maxEndingHere := nums[0]
	maxSoFar := nums[0]
	start, end := 0, 0
	tempStart := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > maxEndingHere+nums[i] {
			tempStart = i
			maxEndingHere = nums[i]
		} else {
			maxEndingHere = maxEndingHere + nums[i]
		}

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
			start = tempStart
			end = i
		}
	}

	return maxSoFar, start, end
}

// DP version
func MaxSubarraySumDyn(nums []int) (int, int, int) {
	if len(nums) == 0 {
		return 0, -1, -1
	}

	maxEndingHere := make([]int, len(nums))
	maxEndingHere[0] = nums[0]
	start, end := 0, 0
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere[i] = max(nums[i], maxEndingHere[i-1]+nums[i])

		if maxSoFar < maxEndingHere[i] {
			maxSoFar = maxEndingHere[i]
			end = i
		}
	}

	// Find the start index
	for i := end; i >= 0; i-- {
		if maxEndingHere[i] == nums[i] {
			start = i
			break
		}
	}

	return maxSoFar, start, end
}

package arrays

// brute approach
// Time Complexity: O(N^2). There are two nested loops traversing the array.
// Space Complexity: O(1).
func MaxWaterBrute(arr []int) int {
	res := 0

	for i, el := range arr {
		left := el
		for j := 0; j < i; j++ {
			left = max(left, arr[j])
		}

		right := el
		for j := i + 1; j < len(arr); j++ {
			right = max(right, arr[j])
		}

		res = res + (min(left, right) - el)
	}

	return res
}

// Time Complexity: O(N)
// Auxiliary Space: O(1)
func MaxWater(arr []int) int {
	//  Indices
	left := 0
	right := len(arr) - 1

	// max
	lMax := 0
	rMax := 0

	res := 0

	for left <= right {
		if rMax <= lMax {
			res += max(0, rMax-arr[right])

			rMax = max(rMax, arr[right])

			right--
			continue
		}
		res += max(0, lMax-arr[left])

		lMax = max(lMax, arr[left])
		left++
	}

	return res
}

// linear search
// Time Complexity: O(N)
// Auxiliary Space: O(1)
func MaxWaterLinear(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	prev := arr[0]

	prevIndex := 0
	res := 0

	temp := 0

	for i := 1; i <= len(arr)-1; i++ {
		if arr[i] >= prev {
			prev = arr[i]
			prevIndex = i

			temp = 0
			continue
		}
		res += prev - arr[i]
		temp += prev - arr[i]

	}

	if prevIndex < len(arr)-1 {
		res -= temp

		prev = arr[len(arr)-1]

		for i := len(arr) - 1; i >= prevIndex; i-- {
			if arr[i] >= prev {
				prev = arr[i]
				continue
			}
			res += prev - arr[i]
		}
	}

	return res
}

package arrays

// or how many platforms we need for trains

// Sweep Line Algorithm
// Time Complexity: O(n), where n is length of time arrays
// Auxiliary space: O(maxDepartureTime)
func OptimizePlaces(arrivals, departures []int) int {
	n := len(arrivals)
	maxDepartureTime := getMax(departures)

	v := make([]int, maxDepartureTime+2)

	for i := 0; i < n; i++ {
		v[arrivals[i]]++
		v[departures[i]+1]--
	}

	count := 0
	maxPlatforms := 0

	for i := 0; i < maxDepartureTime+2; i++ {
		count += v[i]
		maxPlatforms = max(maxPlatforms, count)
	}

	return maxPlatforms
}

func getMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

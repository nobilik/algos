package arrays

// it gives us ordered slice of leaders in same order they appears in parent slice
// Time complexity: O(N)
// Auxiliary space: O(k) // leaders number k < N
// for reduce Auxiliary space we can directelly print output
func OrderedLeaders(arr []int) []int {
	size := len(arr)
	if size == 0 {
		return []int{}
	}
	var maxLeader int
	var leaders []int

	for i := size - 1; i >= 0; i-- {
		if arr[i] > maxLeader {
			maxLeader = arr[i]
			leaders = append([]int{maxLeader}, leaders...)
		}
	}

	return leaders
}

package arrays

// Time Complexity: O(N)
// Auxiliary Space: O(N)
// rotates elementh of array
func RightRotate(a []int, k int) []int {
	n := len(a)
	k = k % n
	rotated := make([]int, n)

	for i := 0; i < n; i++ {
		if i < k {
			// Store rightmost kth elements
			rotated[i] = a[n+i-k]
		} else {
			// Store array after 'k' elements
			rotated[i] = a[i-k]
		}
	}

	return rotated
}

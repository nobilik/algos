package arrays

// Finds missing number between [1,N] in array
// Time Complexity: O(N)
// Auxiliary Space: O(N)
func FindMissingNumberHash(arr []int, N int) int {
	hashMap := make(map[int]bool, N)
	for _, num := range arr {
		hashMap[num] = true
	}

	for i := 1; i <= N; i++ {
		if !hashMap[i] {
			return i
		}
	}

	return 0
}

// Finds missing number between [1,N] in array
// Time Complexity: O(N)
// Auxiliary Space: O(1)
func FindMissingNumber(arr []int, N int) int {
	// Calculate the expected sum of numbers from 1 to N
	// overflow protection to int64
	expectedSum := int64(N) * int64(N+1) / 2

	actualSum := 0
	for _, num := range arr {
		actualSum += num
	}

	// The missing number is the difference between the expected and actual sums
	return int(expectedSum - int64(actualSum))
}

// Finds missing number between [1,N] in array
// Time Complexity: O(N)
// Auxiliary Space: O(1)
func FindMissingNumberXOR(arr []int, N int) int {
	xorSum := 0

	// XOR all numbers from 1 to N
	for i := 1; i <= N; i++ {
		xorSum ^= i
	}

	// XOR all numbers in the array
	for _, num := range arr {
		xorSum ^= num
	}

	return xorSum
}

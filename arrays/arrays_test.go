package arrays

import (
	"math"
	"reflect"
	"testing"
)

func TestMaxSubarraySumSimple(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{name: "Empty array", arr: []int{}, want: 0}, // Edge case: empty array
		{name: "Single element array", arr: []int{5}, want: 5},
		{name: "Positive subarray sum", arr: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{name: "All negative elements", arr: []int{-1, -2, -3, -4}, want: -1},
		{name: "Large positive and negative mix", arr: []int{8, -19, 5, -4, 20, 13}, want: 34},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxSubarraySumSimple(tc.arr)
			if got != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, got)
			}
		})
	}
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums          []int
		expectedSum   int
		expectedStart int
		expectedEnd   int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6, 3, 6},
		{[]int{-2, -3, 4, -1, -2, 1, 5, -3}, 7, 2, 6},
		{[]int{1}, 1, 0, 0},  // Edge case: array of length 1
		{[]int{}, 0, -1, -1}, // Edge case: empty array
	}

	for _, test := range tests {
		maxSum, start, end := MaxSubarraySum(test.nums)
		if maxSum != test.expectedSum || start != test.expectedStart || end != test.expectedEnd {
			t.Errorf("For nums=%v, expected sum=%d, start=%d, end=%d; got sum=%d, start=%d, end=%d",
				test.nums, test.expectedSum, test.expectedStart, test.expectedEnd, maxSum, start, end)
		}
	}
}

func TestFindMissingNumber(t *testing.T) {
	tests := []struct {
		arr      []int
		N        int
		expected int
	}{
		{[]int{1, 2, 4, 5}, 5, 3},    // Missing number is 3
		{[]int{1, 2, 3, 4, 6}, 6, 5}, // Missing number is 5
		{[]int{2, 3, 4, 5}, 5, 1},    // Missing number is 1
		{[]int{1, 2, 3, 4, 5}, 6, 6}, // Missing number is 6
		{[]int{1}, 1, 0},             // Edge case: Missing number is 0
	}

	for _, test := range tests {
		result := FindMissingNumberXOR(test.arr, test.N)
		if result != test.expected {
			t.Errorf("For arr=%v, N=%d, expected=%d, got=%d", test.arr, test.N, test.expected, result)
		}
	}
}

func TestMaxWater(t *testing.T) {
	tests := []struct {
		arr []int

		expected int
	}{
		{[]int{1, 2, 1, 3, 4, 0, 0, 2, 1, 3}, 10}, // Example test case
		{[]int{3, 1, 2, 4, 0, 5}, 7},              // Another test case
		{[]int{1, 1, 1, 1}, 0},                    // All elements are the same
		{[]int{1}, 0},                             // Single element array
		{[]int{}, 0},                              // Empty array
	}

	for _, test := range tests {
		result := MaxWaterLinear(test.arr)
		if result != test.expected {
			t.Errorf("For arr=%v, expected=%d, got=%d", test.arr, test.expected, result)
		}
	}
}

func TestMaxSubarrayProduct(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},     // Example test case: [2,3,-2,4] => 6
		{[]int{-2, 0, -1}, 0},       // Test case with zero: [-2,0,-1] => 0
		{[]int{-2, -3, -1}, 6},      // Test case with all negative numbers: [-2,-3,-1] => 6
		{[]int{1, 2, 3, 4, -5}, 24}, // Test case with both positive and negative numbers: [1,2,3,4,-5] => 24
		{[]int{-1, -2, -3, -4}, 24}, // Test case with all negative numbers: [-1,-2,-3,-4] => 24
		{[]int{0}, 0},               // Test case with single element: [0] => 0
		{[]int{}, math.MinInt64},    // Test case with empty array: [] => MinInt64
	}

	for _, test := range tests {
		result := MaxSubarrayProduct(test.arr)
		if result != test.expected {
			t.Errorf("For arr=%v, expected=%d, got=%d", test.arr, test.expected, result)
		}
	}
}

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},    // Example test case: [1, 7, 3, 6, 5, 6] => 3
		{[]int{1, 2, 3}, -1},            // Test case without pivot: [1, 2, 3] => -1
		{[]int{2, 1, -1}, 0},            // Test case with negative numbers: [2, 1, -1] => 0
		{[]int{-1, -1, -1, 0, 1, 1}, 0}, // Test case with negative and positive numbers: [-1, -1, -1, 0, 1, 1] => 0
		{[]int{1}, 0},                   // Test case with single element: [1] => 0
		{[]int{}, -1},                   // Test case with empty array: [] => -1
	}

	for _, test := range tests {
		result := EquilibriumIndex(test.nums)
		if result != test.expected {
			t.Errorf("For nums=%v, expected=%d, got=%d", test.nums, test.expected, result)
		}
	}
}

func TestOrderedLeaders(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{16, 17, 4, 3, 5, 2}, []int{17, 5, 2}}, // Example test case: [16, 17, 4, 3, 5, 2] => [17, 5, 2]
		{[]int{1, 2, 3, 4, 5}, []int{5}},             // Test case with all elements in ascending order: [1, 2, 3, 4, 5] => [5]
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}}, // Test case with all elements in descending order: [5, 4, 3, 2, 1] => [5, 4, 3, 2, 1]
		{[]int{5, 5, 5, 5, 5}, []int{5}},             // Test case with all elements equal: [5, 5, 5, 5, 5] => [5]
		{[]int{}, []int{}},                           // Test case with empty array: [] => []
	}

	for _, test := range tests {
		result := OrderedLeaders(test.arr)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For arr=%v, expected=%v, got=%v", test.arr, test.expected, result)
		}
	}
}

func TestFindPlatformOptimized(t *testing.T) {
	tests := []struct {
		arr      []int
		dep      []int
		expected int
	}{
		{[]int{900, 940, 950, 1100, 1500, 1800}, []int{910, 1200, 1120, 1130, 1900, 2000}, 3}, // Example test case
		{[]int{200, 210, 300, 320, 500}, []int{230, 320, 400, 430, 600}, 3},                   // Another test case
		{[]int{100, 200, 300}, []int{150, 250, 350}, 1},                                       // Test case with one platform needed
	}

	for _, test := range tests {
		result := OptimizePlaces(test.arr, test.dep)
		if result != test.expected {
			t.Errorf("For arr=%v, dep=%v, expected=%d, got=%d", test.arr, test.dep, test.expected, result)
		}
	}
}

func TestRightRotate(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}}, // Example test case
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}}, // Test case with larger rotation count
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}}, // Test case with rotation count 0
	}

	for _, test := range tests {
		result := RightRotate(test.arr, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For arr=%v, k=%d, expected=%v, got=%v", test.arr, test.k, test.expected, result)
		}
	}
}

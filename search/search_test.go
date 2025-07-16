package search

import "testing"

func TestBinarySearch(t *testing.T) {
	// Test cases for BinarySearch
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected bool
	}{
		{
			name:     "Find element in middle",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: true,
		},
		{
			name:     "Find first element",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: true,
		},
		{
			name:     "Find last element",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: true,
		},
		{
			name:     "Element not found - too small",
			arr:      []int{1, 2, 3, 4, 5},
			target:   0,
			expected: false,
		},
		{
			name:     "Element not found - too large",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: false,
		},
		{
			name:     "Element not found - in range",
			arr:      []int{1, 3, 5, 7, 9},
			target:   4,
			expected: false,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			target:   1,
			expected: false,
		},
		{
			name:     "Single element - found",
			arr:      []int{42},
			target:   42,
			expected: true,
		},
		{
			name:     "Single element - not found",
			arr:      []int{42},
			target:   24,
			expected: false,
		},
		{
			name:     "Larger sorted array",
			arr:      []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target:   70,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %v, expected %v", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestSequentialSearch(t *testing.T) {
	// Test cases for SequentialSearch
	tests := []struct {
		name     string
		arr      []int
		value    int
		expected bool
	}{
		{
			name:     "Find element at beginning",
			arr:      []int{1, 2, 3, 4, 5},
			value:    1,
			expected: true,
		},
		{
			name:     "Find element at end",
			arr:      []int{1, 2, 3, 4, 5},
			value:    5,
			expected: true,
		},
		{
			name:     "Find element in middle",
			arr:      []int{1, 2, 3, 4, 5},
			value:    3,
			expected: true,
		},
		{
			name:     "Element not found",
			arr:      []int{1, 2, 3, 4, 5},
			value:    6,
			expected: false,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			value:    1,
			expected: false,
		},
		{
			name:     "Single element - found",
			arr:      []int{42},
			value:    42,
			expected: true,
		},
		{
			name:     "Single element - not found",
			arr:      []int{42},
			value:    24,
			expected: false,
		},
		{
			name:     "Duplicate elements",
			arr:      []int{1, 2, 2, 3, 4},
			value:    2,
			expected: true,
		},
		{
			name:     "Unsorted array",
			arr:      []int{5, 2, 8, 1, 9},
			value:    8,
			expected: true,
		},
		{
			name:     "Negative numbers",
			arr:      []int{-3, -1, 0, 2, 5},
			value:    -1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SequentialSearch(tt.arr, tt.value)
			if result != tt.expected {
				t.Errorf("SequentialSearch(%v, %d) = %v, expected %v", tt.arr, tt.value, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkBinarySearch(b *testing.B) {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, 500)
	}
}

func BenchmarkSequentialSearch(b *testing.B) {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SequentialSearch(arr, 500)
	}
}
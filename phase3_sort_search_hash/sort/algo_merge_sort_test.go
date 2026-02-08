package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	is := assert.New(t)

	t.Run("MergerSort", func(t *testing.T) {
		t.Run("GeneralUnsortedArray", func(t *testing.T) {
			// Test with a typical unsorted list
			input := []int{38, 27, 43, 3, 9, 82, 10}
			expected := []int{3, 9, 10, 27, 38, 43, 82}
			is.Equal(expected, MergeSort(input))
		})

		t.Run("LargeRandomArray", func(t *testing.T) {
			// Testing O(n log n) efficiency with more elements
			input := []int{5, 2, 9, 1, 5, 6, 3, 0, -1, 8, 7, 4}
			expected := []int{-1, 0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
			is.Equal(expected, MergeSort(input))
		})

		t.Run("AlreadySorted", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			expected := []int{1, 2, 3, 4, 5}
			is.Equal(expected, MergeSort(input))
		})

		t.Run("ReverseSorted", func(t *testing.T) {
			input := []int{5, 4, 3, 2, 1}
			expected := []int{1, 2, 3, 4, 5}
			is.Equal(expected, MergeSort(input))
		})

		t.Run("NegativeNumbers", func(t *testing.T) {
			input := []int{-5, -1, -10, 0, 5, 2}
			expected := []int{-10, -5, -1, 0, 2, 5}
			is.Equal(expected, MergeSort(input))
		})

		t.Run("EdgeCases", func(t *testing.T) {
			// Empty array
			is.Equal([]int{}, MergeSort([]int{}))
			// Single element
			is.Equal([]int{1}, MergeSort([]int{1}))
		})
	})
}

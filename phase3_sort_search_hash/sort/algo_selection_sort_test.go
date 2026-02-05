package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	is := assert.New(t)
	// Assuming SelectionSort is in the same package or imported

	t.Run("InsertionSort", func(t *testing.T) {
		t.Run("UnsortedArray", func(t *testing.T) {
			// Normal case with random numbers
			input := []int{29, 10, 14, 37, 13}
			expected := []int{10, 13, 14, 29, 37}
			is.Equal(expected, SelectionSort(input))
		})

		t.Run("AlreadySorted", func(t *testing.T) {
			// Should remain unchanged
			input := []int{1, 2, 3, 4, 5}
			expected := []int{1, 2, 3, 4, 5}
			is.Equal(expected, SelectionSort(input))
		})

		t.Run("ReverseSorted", func(t *testing.T) {
			// Worst case for moves
			input := []int{5, 4, 3, 2, 1}
			expected := []int{1, 2, 3, 4, 5}
			is.Equal(expected, SelectionSort(input))
		})

		t.Run("WithDuplicates", func(t *testing.T) {
			// Handling identical values
			input := []int{3, 1, 2, 1, 3}
			expected := []int{1, 1, 2, 3, 3}
			is.Equal(expected, SelectionSort(input))
		})

		t.Run("EmptyAndSingleElement", func(t *testing.T) {
			is.Equal([]int{}, SelectionSort([]int{}))
			is.Equal([]int{1}, SelectionSort([]int{1}))
		})
	})
}

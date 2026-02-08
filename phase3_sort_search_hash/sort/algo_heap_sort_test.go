package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	is := assert.New(t)

	t.Run("TestHeapSort", func(t *testing.T) {
		t.Run("Standard_Unsorted_Array", func(t *testing.T) {
			input := []int{12, 11, 13, 5, 6, 7}
			expected := []int{5, 6, 7, 11, 12, 13}
			HeapSort(input)
			is.Equal(expected, input)
		})

		t.Run("Already_Sorted", func(t *testing.T) {
			input := []int{1, 2, 3, 4, 5}
			expected := []int{1, 2, 3, 4, 5}
			HeapSort(input)
			is.Equal(expected, input)
		})

		t.Run("Reverse_Sorted", func(t *testing.T) {
			input := []int{10, 8, 6, 4, 2}
			expected := []int{2, 4, 6, 8, 10}
			HeapSort(input)
			is.Equal(expected, input)
		})

		t.Run("Array_With_Duplicates", func(t *testing.T) {
			input := []int{4, 10, 3, 5, 1, 4, 1}
			expected := []int{1, 1, 3, 4, 4, 5, 10}
			HeapSort(input)
			is.Equal(expected, input)
		})

		t.Run("Negative_Numbers", func(t *testing.T) {
			input := []int{-1, -5, 2, 0, -2}
			expected := []int{-5, -2, -1, 0, 2}
			HeapSort(input)
			is.Equal(expected, input)
		})

		t.Run("Single_And_Empty", func(t *testing.T) {
			is.NotPanics(func() {
				HeapSort([]int{42})
				HeapSort([]int{})
			})
		})
	})
}

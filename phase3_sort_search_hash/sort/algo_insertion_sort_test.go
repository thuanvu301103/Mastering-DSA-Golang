package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	is := assert.New(t)

	t.Run("InsertionSort", func(t *testing.T) {
		t.Run("RandomOrder", func(t *testing.T) {
			nums := []int{12, 11, 13, 5, 6}
			expected := []int{5, 6, 11, 12, 13}
			InsertionSort(nums)
			is.Equal(expected, nums)
		})

		t.Run("NearlySorted", func(t *testing.T) {
			// Insertion sort is very fast for this case
			nums := []int{1, 2, 3, 5, 4}
			expected := []int{1, 2, 3, 4, 5}
			InsertionSort(nums)
			is.Equal(expected, nums)
		})

		t.Run("ReverseOrder", func(t *testing.T) {
			nums := []int{5, 4, 3, 2, 1}
			expected := []int{1, 2, 3, 4, 5}
			InsertionSort(nums)
			is.Equal(expected, nums)
		})

		t.Run("EmptyAndSingle", func(t *testing.T) {
			empty := []int{}
			InsertionSort(empty)
			is.Equal([]int{}, empty)

			single := []int{1}
			InsertionSort(single)
			is.Equal([]int{1}, single)
		})
	})
}

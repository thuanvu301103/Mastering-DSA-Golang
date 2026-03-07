package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	is := assert.New(t)

	t.Run("Element_Exists_In_Middle", func(t *testing.T) {
		nums := []int{10, 20, 30, 40, 50}
		target := 30

		index := LinearSearch(nums, target)
		is.Equal(2, index, "Should find 30 at index 2")
	})

	t.Run("Element_At_Start", func(t *testing.T) {
		nums := []int{99, 50, 20}
		is.Equal(0, LinearSearch(nums, 99))
	})

	t.Run("Element_At_End", func(t *testing.T) {
		nums := []int{5, 10, 15}
		is.Equal(2, LinearSearch(nums, 15))
	})

	t.Run("Element_Does_Not_Exist", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		is.Equal(-1, LinearSearch(nums, 100), "Should return -1 for missing element")
	})

	t.Run("Empty_Slice", func(t *testing.T) {
		var nums []int
		is.Equal(-1, LinearSearch(nums, 5), "Searching in empty slice should return -1")
	})

	t.Run("Duplicate_Elements", func(t *testing.T) {
		nums := []int{7, 8, 7, 9}
		// Linear search usually returns the FIRST occurrence
		is.Equal(0, LinearSearch(nums, 7), "Should return the index of the first occurrence")
	})
}

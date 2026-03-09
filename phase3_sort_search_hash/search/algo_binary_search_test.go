package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	is := assert.New(t)

	t.Run("Target_In_Middle", func(t *testing.T) {
		nums := []int{10, 20, 30, 40, 50} // Odd length
		is.Equal(2, BinarySearch(nums, 30))
	})

	t.Run("Target_At_Start", func(t *testing.T) {
		nums := []int{10, 20, 30, 40} // Even length
		is.Equal(0, BinarySearch(nums, 10))
	})

	t.Run("Target_At_End", func(t *testing.T) {
		nums := []int{10, 20, 30, 40}
		is.Equal(3, BinarySearch(nums, 40))
	})

	t.Run("Target_Not_Found", func(t *testing.T) {
		nums := []int{1, 3, 5, 7, 9}
		is.Equal(-1, BinarySearch(nums, 4), "Should return -1 for missing element")
	})

	t.Run("Target_Smaller_Than_All", func(t *testing.T) {
		nums := []int{10, 20, 30}
		is.Equal(-1, BinarySearch(nums, 5))
	})

	t.Run("Target_Larger_Than_All", func(t *testing.T) {
		nums := []int{10, 20, 30}
		is.Equal(-1, BinarySearch(nums, 100))
	})

	t.Run("Empty_Slice", func(t *testing.T) {
		var nums []int
		is.Equal(-1, BinarySearch(nums, 1))
	})

	t.Run("Single_Element_Found", func(t *testing.T) {
		nums := []int{42}
		is.Equal(0, BinarySearch(nums, 42))
	})
}

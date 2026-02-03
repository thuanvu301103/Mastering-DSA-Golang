package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoPointerTechniques(t *testing.T) {
	is := assert.New(t)
	ops := TwoPointerOps{}

	// --- Strategy 1: Left/Right Pointers (Convergence) ---
	// Logic: Pointers start at boundaries and move toward each other.
	t.Run("LeftRight_Convergence", func(t *testing.T) {

		t.Run("TwoSumSorted_Found", func(t *testing.T) {
			// Tests O(n) search on sorted data
			nums := []int{2, 7, 11, 15}
			target := 9
			// Should find indices 0 and 1
			is.Equal([]int{0, 1}, ops.TwoSumSorted(nums, target))
		})

		t.Run("IsPalindrome_MixedCase", func(t *testing.T) {
			// Tests pointers meeting in the exact middle
			is.True(ops.IsPalindrome("Racecar"))
			is.False(ops.IsPalindrome("Golang"))
		})

		t.Run("ReverseArray_InPlace", func(t *testing.T) {
			// Tests swapping elements using O(1) extra space
			arr := []int{1, 2, 3, 4, 5}
			ops.Reverse(arr)
			is.Equal([]int{5, 4, 3, 2, 1}, arr)
		})
	})
	/*
		// --- Strategy 2: Fast/Slow Pointers (Distance) ---
		// Logic: One pointer moves at 2x speed to find cycles or midpoints.
		t.Run("FastSlow_Distance", func(t *testing.T) {

			t.Run("DetectCycle_Positive", func(t *testing.T) {
				// Tests if 'fast' laps 'slow' in a circular structure
				list := createCyclicList([]int{10, 20, 30, 40}, 1) // 40 points back to 20
				is.True(ops.HasCycle(list))
			})
			/*
				t.Run("DetectCycle_Negative", func(t *testing.T) {
					list := createLinkedList([]int{1, 2, 3})
					is.False(ops.HasCycle(list))
				})

				t.Run("FindMiddle_EvenLength", func(t *testing.T) {
					// In [1,2,3,4], usually returns the second middle (3)
					list := createLinkedList([]int{1, 2, 3, 4})
					middle := ops.FindMiddle(list)
					is.Equal(3, middle.Value)
				})
		})
			// --- Strategy 3: Edge Cases ---
			t.Run("BoundaryConditions", func(t *testing.T) {
				is.Empty(ops.TwoSumSorted([]int{1, 5}, 10), "Should handle no match")
				is.True(ops.IsPalindrome("a"), "Single character is a palindrome")
				is.Nil(ops.FindMiddle(nil), "Should handle empty linked list")
			})
	*/
}

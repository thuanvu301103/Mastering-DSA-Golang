package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlidingWindow(t *testing.T) {
	is := assert.New(t)
	ops := SlidingWindowOps{}

	// --- SECTION 1: FIXED WINDOW ---
	// Logic: Size K is constant. Add incoming, remove outgoing.
	t.Run("FixedWindow", func(t *testing.T) {

		t.Run("MaxSumSubarray_SizeK", func(t *testing.T) {
			nums := []int{2, 1, 5, 1, 3, 2}
			k := 3
			// [5, 1, 3] gives the max sum of 9
			is.Equal(9, ops.MaxSumFixed(nums, k))
		})

		t.Run("AveragesOfSubarrays", func(t *testing.T) {
			nums := []int{1, 3, 2, 6, -1}
			k := 5
			// Average of all 5 elements is 11/5 = 2.2
			expected := []float64{2.2}
			is.Equal(expected, ops.FindAverages(nums, k))
		})
	})

	// --- SECTION 2: VARIABLE WINDOW (DYNAMIC) ---
	// Logic: Window size changes based on a condition (Stretch & Shrink).
	t.Run("VariableWindow", func(t *testing.T) {

		t.Run("SmallestSubarrayWithSum", func(t *testing.T) {
			targetSum := 7
			nums := []int{2, 1, 5, 2, 3, 2}
			// Smallest subarray with sum >= 7 is [5, 2], length = 2
			is.Equal(2, ops.MinSizeSubarraySum(targetSum, nums))
		})

		t.Run("LongestSubstringKDistinct", func(t *testing.T) {
			str := "araaci"
			k := 2
			// Longest substring with 2 distinct characters is "araa", length = 4
			is.Equal(4, ops.LongestSubstringKDistinct(str, k))
		})
	})

	// --- SECTION 3: EDGE CASES ---
	t.Run("EdgeCases", func(t *testing.T) {
		is.Equal(0, ops.MaxSumFixed([]int{1, 2}, 5), "K larger than array size")
		is.Equal(0, ops.MinSizeSubarraySum(100, []int{1, 2, 3}), "Sum never reached")
	})

}

package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixSumSuite(t *testing.T) {
	is := assert.New(t)
	ops := PrefixSumOps{}

	// --- SECTION 1: 1D PREFIX SUM ---
	t.Run("1D_PrefixSum", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6}
		// Preprocessing: Build the prefix sum array
		// P = [1, 3, 6, 10, 15, 21]
		p := ops.BuildPrefixArray(nums)

		t.Run("RangeSumQuery", func(t *testing.T) {
			// Sum from index 2 to 4 (values: 3+4+5 = 12)
			// Logic: P[4] - P[1] = 15 - 3 = 12
			is.Equal(12, ops.GetRangeSum(p, 2, 4))
		})

		t.Run("FullRange", func(t *testing.T) {
			is.Equal(21, ops.GetRangeSum(p, 0, 5))
		})
	})

	// --- SECTION 2: 2D PREFIX SUM (MATRIX) ---
	t.Run("2D_PrefixSum", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		// Preprocessing: Build the 2D prefix matrix
		pMatrix := ops.Build2DPrefixMatrix(matrix)

		t.Run("SubRectangleSum", func(t *testing.T) {
			// Rectangle from (1,1) to (2,2)
			// Values: 5, 6, 8, 9 -> Sum = 28
			is.Equal(28, ops.GetRectSum(pMatrix, 1, 1, 2, 2))
		})

		t.Run("SingleCell", func(t *testing.T) {
			is.Equal(1, ops.GetRectSum(pMatrix, 0, 0, 0, 0))
		})
	})
}

package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArrayOps(t *testing.T) {
	is := assert.New(t)
	ops := DynamicArrayOps{}

	t.Run("AppendAndLength", func(t *testing.T) {
		// Testing how slices grow in length
		initial := []int{1, 2, 3}
		result := ops.AppendElement(initial, 4)

		is.Equal([]int{1, 2, 3, 4}, result)
		is.Equal(4, len(result))
	})

	t.Run("FilterElements", func(t *testing.T) {
		// Testing creating a smaller slice from a larger one based on criteria
		input := []int{1, 2, 3, 4, 5, 6}
		expected := []int{2, 4, 6}
		is.Equal(expected, ops.FilterEvenNumbers(input))
	})

	t.Run("SubSlicing", func(t *testing.T) {
		// Testing the [low:high] operator logic
		input := []int{10, 20, 30, 40, 50}
		// Expected to take indices 1, 2, and 3
		result := ops.GetRange(input, 1, 4)
		is.Equal([]int{20, 30, 40}, result)
	})

	t.Run("CapacityVsLength", func(t *testing.T) {
		// Testing internal re-allocation when capacity is exceeded
		slice := make([]int, 0, 2) // length 0, capacity 2
		slice = append(slice, 1, 2)
		oldCap := cap(slice)

		// This append forces the slice to allocate a new, larger underlying array
		slice = append(slice, 3)
		is.Greater(cap(slice), oldCap, "Capacity should double or increase during re-allocation")
	})

	t.Run("ReferenceBehavior", func(t *testing.T) {
		// Testing that slices pass a pointer to the same underlying array
		original := []int{1, 2, 3}
		ops.ModifyElement(original, 0, 99)

		is.Equal(99, original[0], "Changes inside the function should reflect in the original slice")
	})
}

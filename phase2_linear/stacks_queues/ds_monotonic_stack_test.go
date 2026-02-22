package stack_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonotonicStack(t *testing.T) {
	is := assert.New(t)

	t.Run("Monotonic_Increasing_Logic", func(t *testing.T) {
		// Goal: Maintain elements in increasing order (e.g., 1, 3, 5)
		// If we push '2', it must pop '3' and '5' to keep the order [1, 2].
		ms := &MonotonicStack{Type: "increasing"}

		ms.Push(1)
		ms.Push(3)
		ms.Push(5)
		ms.Push(2) // 3 and 5 are popped because they are > 2

		is.Equal([]int{1, 2}, ms.GetData(), "Stack should maintain [1, 2]")
	})

	t.Run("Monotonic_Decreasing_Logic", func(t *testing.T) {
		// Goal: Maintain elements in decreasing order (e.g., 10, 8, 6)
		// If we push '9', it must pop '8' and '6' to keep the order [10, 9].
		ms := &MonotonicStack{Type: "decreasing"}

		ms.Push(10)
		ms.Push(8)
		ms.Push(6)
		ms.Push(9) // 8 and 6 are popped because they are < 9

		is.Equal([]int{10, 9}, ms.GetData(), "Stack should maintain [10, 9]")
	})

	t.Run("Daily_Temperatures_Real_World_Case", func(t *testing.T) {
		// Use Case: Next Greater Element (Decreasing Stack)
		// Input: [73, 74, 75, 71, 69, 72, 76, 73]
		// Output: Number of days until a warmer temperature.
		input := []int{73, 74, 75, 71, 69, 72, 76, 73}
		expected := []int{1, 1, 4, 2, 1, 1, 0, 0}

		is.Equal(expected, SolveDailyTemperatures(input))
	})
}

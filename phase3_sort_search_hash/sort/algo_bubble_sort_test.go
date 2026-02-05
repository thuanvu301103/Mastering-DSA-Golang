package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	is := assert.New(t)

	t.Run("BubbleSort", func(t *testing.T) {

		t.Run("StandardCase", func(t *testing.T) {
			nums := []int{64, 34, 25, 12, 22, 11, 90}
			expected := []int{11, 12, 22, 25, 34, 64, 90}

			BubbleSort(nums)
			is.Equal(expected, nums)
		})

		t.Run("WorstCase_Reversed", func(t *testing.T) {
			nums := []int{5, 4, 3, 2, 1}
			expected := []int{1, 2, 3, 4, 5}

			BubbleSort(nums)
			is.Equal(expected, nums)
		})

		t.Run("BestCase_AlreadySorted", func(t *testing.T) {
			nums := []int{1, 2, 3, 4, 5}
			expected := []int{1, 2, 3, 4, 5}

			BubbleSort(nums)
			is.Equal(expected, nums)
		})

		t.Run("Duplicates", func(t *testing.T) {
			nums := []int{3, 1, 2, 3, 1}
			expected := []int{1, 1, 2, 3, 3}

			BubbleSort(nums)
			is.Equal(expected, nums)
		})

		t.Run("EdgeCases", func(t *testing.T) {
			empty := []int{}
			BubbleSort(empty)
			is.Equal([]int{}, empty)

			single := []int{1}
			BubbleSort(single)
			is.Equal([]int{1}, single)
		})
	})
}

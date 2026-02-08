package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	is := assert.New(t)

	t.Run("Standard_Case", func(t *testing.T) {
		input := []int{10, 7, 8, 9, 1, 5}
		expected := []int{1, 5, 7, 8, 9, 10}

		// Much cleaner: no more 0 and len-1 here!
		QuickSort(input)

		is.Equal(expected, input)
	})

	t.Run("Empty_Array", func(t *testing.T) {
		input := []int{}
		QuickSort(input)
		is.Equal([]int{}, input)
	})
}

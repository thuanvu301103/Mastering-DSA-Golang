package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	is := assert.New(t)

	t.Run("TextMaxHeap", func(t *testing.T) {
		t.Run("Insert_Elements", func(t *testing.T) {
			h := &MaxHeap{}
			elements := []int{10, 20, 5, 30, 15}

			for _, el := range elements {
				h.Insert(el)
			}

			// The root must be the maximum value (30)
			is.Equal(30, h.Peek(), "Root should be the maximum element")
		})

		t.Run("ExtractMax_Order", func(t *testing.T) {
			h := &MaxHeap{}
			h.Insert(50)
			h.Insert(100)
			h.Insert(30)
			h.Insert(80)

			// Should extract in descending order: 100 -> 80 -> 50 -> 30
			is.Equal(100, h.ExtractMax())
			is.Equal(80, h.ExtractMax())
			is.Equal(50, h.ExtractMax())
			is.Equal(30, h.ExtractMax())
		})

		t.Run("Heapify_Property", func(t *testing.T) {
			// Testing internal array structure
			h := &MaxHeap{}
			nums := []int{1, 2, 3}
			// After inserting 1, 2, 3:
			//   1 -> Insert 2 (swap with 1) -> 2, 1 -> Insert 3 (swap with 2) -> 3, 1, 2
			for _, n := range nums {
				h.Insert(n)
			}

			// Structure should follow: [3, 1, 2]
			// Parent (3) > Left (1) and Right (2)
			is.Equal([]int{3, 1, 2}, h.GetArray())
		})

		t.Run("Empty_Heap", func(t *testing.T) {
			h := &MaxHeap{}
			is.Panics(func() { h.Peek() }, "Peeking an empty heap should handle error or panic")
		})
	})
}

package tree

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	is := assert.New(t)

	t.Run("Insert_Maintains_Max_Property", func(t *testing.T) {
		h := &MaxHeap{}
		// Inserting out of order
		h.Insert(10)
		h.Insert(30)
		h.Insert(20)
		h.Insert(5)

		// The root (index 0) must always be the maximum value
		is.Equal(30, h.Peek(), "Root should be 30 after insertions")
	})

	t.Run("ExtractMax_Sequential_Order", func(t *testing.T) {
		h := &MaxHeap{}
		h.Insert(50)
		h.Insert(100)
		h.Insert(75)
		h.Insert(25)

		// Should extract elements in descending order
		is.Equal(100, h.ExtractMax())
		is.Equal(75, h.ExtractMax())
		is.Equal(50, h.ExtractMax())
		is.Equal(25, h.ExtractMax())
		
		is.True(h.IsEmpty())
	})

	t.Run("Heapify_Array_Structure", func(t *testing.T) {
		// Verifying the internal array representation
		// Insert: 10, 20, 15
		// Step 1: [10]
		// Step 2: [20, 10] (20 bubbles up)
		// Step 3: [20, 10, 15]
		h := &MaxHeap{}
		h.Insert(10)
		h.Insert(20)
		h.Insert(15)

		expectedArray := []int{20, 10, 15}
		is.Equal(expectedArray, h.GetInternalSlice())
	})

	t.Run("Empty_Heap_Edge_Cases", func(t *testing.T) {
		h := &MaxHeap{}
		
		_, ok := h.ExtractMaxSafe()
		is.False(ok, "Extracting from empty heap should return false/error")
		
		is.Panics(func() { h.Peek() }, "Peeking empty heap should be handled or panic")
	})
}
package stack_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	is := assert.New(t)

	t.Run("Basic_Enqueue_Dequeue", func(t *testing.T) {
		q := NewCircularQueue(3)

		q.Enqueue(1)
		q.Enqueue(2)

		val, _ := q.Dequeue()
		is.Equal(1, val)

		val, _ = q.Dequeue()
		is.Equal(2, val)
	})

	t.Run("Wrap_Around_Logic", func(t *testing.T) {
		// Size is 3.
		q := NewCircularQueue(3)
		q.Enqueue(10)
		q.Enqueue(20)
		q.Enqueue(30)

		// Queue is now full. Remove one to make space at the beginning.
		q.Dequeue() // Removes 10

		// This should wrap around and occupy the index where 10 used to be
		success := q.Enqueue(40)
		is.True(success, "Should be able to enqueue 40 after dequeuing 10")

		// Verify order: 20 -> 30 -> 40
		v1, _ := q.Dequeue()
		v2, _ := q.Dequeue()
		v3, _ := q.Dequeue()

		is.Equal(20, v1)
		is.Equal(30, v2)
		is.Equal(40, v3)
	})

	t.Run("Full_and_Empty_States", func(t *testing.T) {
		q := NewCircularQueue(2)

		is.True(q.IsEmpty())

		q.Enqueue(1)
		q.Enqueue(2)

		is.True(q.IsFull())
		is.False(q.Enqueue(3), "Should return false when enqueuing to a full queue")

		q.Dequeue()
		q.Dequeue()
		is.True(q.IsEmpty())
		is.False(q.IsFull())
	})
}

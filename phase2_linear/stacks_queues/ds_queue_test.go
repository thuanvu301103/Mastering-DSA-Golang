package stack_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	is := assert.New(t)

	t.Run("Enqueue_and_Dequeue_Order", func(t *testing.T) {
		q := &Queue{}

		// Enqueue 10, then 20
		q.Enqueue(10)
		q.Enqueue(20)

		// First in (10) must be first out (FIFO)
		val, ok := q.Dequeue()
		is.True(ok)
		is.Equal(10, val)

		// Next out should be 20
		val, ok = q.Dequeue()
		is.True(ok)
		is.Equal(20, val)
	})

	t.Run("Front_Method", func(t *testing.T) {
		q := &Queue{}
		q.Enqueue(100)
		q.Enqueue(200)

		// Front should show 100 but NOT remove it
		val, ok := q.Front()
		is.True(ok)
		is.Equal(100, val)

		// Verify 100 is still there by calling Front again
		val, _ = q.Front()
		is.Equal(100, val)
	})

	t.Run("IsEmpty_Behavior", func(t *testing.T) {
		q := &Queue{}
		is.True(q.IsEmpty(), "New queue should be empty")

		q.Enqueue(1)
		is.False(q.IsEmpty(), "Queue should not be empty after Enqueue")

		q.Dequeue()
		is.True(q.IsEmpty(), "Queue should be empty after Dequeueing all elements")
	})

	t.Run("Empty_Queue_Edge_Cases", func(t *testing.T) {
		q := &Queue{}

		// Dequeue from empty queue
		val, ok := q.Dequeue()
		is.False(ok)
		is.Equal(0, val)

		// Front from empty queue
		val, ok = q.Front()
		is.False(ok)
		is.Equal(0, val)
	})
}

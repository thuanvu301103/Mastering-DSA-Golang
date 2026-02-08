package stack_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	is := assert.New(t)

	t.Run("Test_Stack", func(t *testing.T) {
		t.Run("Push_and_Pop_Order", func(t *testing.T) {
			s := &Stack{}

			// Push 1, then 2
			s.Push(1)
			s.Push(2)

			// Should pop 2 first (LIFO)
			val, ok := s.Pop()
			is.True(ok)
			is.Equal(2, val)

			// Then pop 1
			val, ok = s.Pop()
			is.True(ok)
			is.Equal(1, val)
		})

		t.Run("Peek_Property", func(t *testing.T) {
			s := &Stack{}
			s.Push(100)

			// Peek should show 100 but not remove it
			val, ok := s.Peek()
			is.True(ok)
			is.Equal(100, val)
			is.Equal(1, len(*s), "Stack size should still be 1 after Peek")
		})

		t.Run("Empty_Stack_Behavior", func(t *testing.T) {
			s := &Stack{}

			// Pop from empty should return false
			val, ok := s.Pop()
			is.False(ok)
			is.Equal(0, val)

			// Peek from empty should return false
			val, ok = s.Peek()
			is.False(ok)
			is.Equal(0, val)
		})

		t.Run("Large_Sequence", func(t *testing.T) {
			s := &Stack{}
			for i := 0; i < 1000; i++ {
				s.Push(i)
			}

			val, _ := s.Peek()
			is.Equal(999, val)
			is.Equal(1000, len(*s))
		})
	})
}

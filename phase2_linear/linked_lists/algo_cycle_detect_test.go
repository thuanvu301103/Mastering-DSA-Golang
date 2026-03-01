package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList_HasCycle(t *testing.T) {
	is := assert.New(t)

	t.Run("Linear_List_No_Cycle", func(t *testing.T) {
		l := &SinglyLinkedList{}
		l.PushBack(1)
		l.PushBack(2)
		l.PushBack(3)

		is.False(l.HasCycle(), "Linear list should not have a cycle")
	})

	t.Run("Circular_List_Detected", func(t *testing.T) {
		l := &SinglyLinkedList{}
		l.PushBack(10)
		l.PushBack(20)
		l.PushBack(30)
		l.PushBack(40)

		// Manually create a cycle: 40 -> 20
		// 10 -> 20 -> 30 -> 40 --|
		//        ^               |
		//        |_______________|

		node40 := l.Head.Next.Next.Next
		node20 := l.Head.Next
		node40.Next = node20

		is.True(l.HasCycle(), "Should detect the cycle where 40 points back to 20")
	})

	t.Run("Single_Node_Loop", func(t *testing.T) {
		l := &SinglyLinkedList{}
		l.PushBack(1)

		// Create a self-loop: 1 -> 1
		l.Head.Next = l.Head

		is.True(l.HasCycle())
	})

	t.Run("Empty_List", func(t *testing.T) {
		l := &SinglyLinkedList{}
		is.False(l.HasCycle())
	})
}

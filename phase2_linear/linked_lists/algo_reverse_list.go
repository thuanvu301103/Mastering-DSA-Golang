package linked_list

import "fmt"

// Push adds a new node at the beginning (Prepend) - O(1)
func (l *SinglyLinkedList) Push(data int) {
	newNode := &Node{Value: data}
	newNode.Next = l.Head
	l.Head = newNode
}

// Reverse implements the three-pointer strategy to reverse the list - O(n)
func (l *SinglyLinkedList) Reverse() {
	var prev *Node = nil
	curr := l.Head
	var next *Node = nil

	for curr != nil {
		// 1. Store the next node to avoid losing the rest of the list
		next = curr.Next

		// 2. Reverse the current node's pointer to face backward
		curr.Next = prev

		// 3. Move 'prev' and 'curr' one step forward
		prev = curr
		curr = next
	}

	// 4. Update the list head to the last node processed (new front)
	l.Head = prev
}

// Display prints the linked list in a readable format
func (l *SinglyLinkedList) Display() {
	if l.Head == nil {
		fmt.Println("List is empty.")
		return
	}

	curr := l.Head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}

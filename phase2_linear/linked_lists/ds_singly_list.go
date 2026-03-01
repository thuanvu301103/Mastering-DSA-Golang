package linked_list

// Node represents a single element in the singly linked list
type Node struct {
	Value int
	Next  *Node
}

// SinglyLinkedList represents the linear data structure with one-way pointers
type SinglyLinkedList struct {
	Head *Node
	Size int
}

// PushFront adds an element to the beginning of the list - O(1)
func (l *SinglyLinkedList) PushFront(val int) {
	newNode := &Node{
		Value: val,
		Next:  l.Head, // Point new node to the current head
	}
	l.Head = newNode
	l.Size++
}

// PushBack adds an element to the end of the list - O(n)
func (l *SinglyLinkedList) PushBack(val int) {
	newNode := &Node{Value: val}
	if l.Head == nil {
		l.Head = newNode
	} else {
		// Traverse to the last node
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.Size++
}

// Remove deletes the first occurrence of a value - O(n)
func (l *SinglyLinkedList) Remove(val int) bool {
	if l.IsEmpty() {
		return false
	}

	// Case 1: Value is at the Head
	if l.Head.Value == val {
		l.Head = l.Head.Next
		l.Size--
		return true
	}

	// Case 2: Search in the rest of the list
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == val {
			// Skip the node to be removed
			current.Next = current.Next.Next
			l.Size--
			return true
		}
		current = current.Next
	}
	return false
}

// Find checks if a value exists in the list - O(n)
func (l *SinglyLinkedList) Find(val int) bool {
	current := l.Head
	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}
	return false
}

// IsEmpty returns true if list has no nodes
func (l *SinglyLinkedList) IsEmpty() bool {
	return l.Size == 0 || l.Head == nil
}

// Clear resets the list
func (l *SinglyLinkedList) Clear() {
	l.Head = nil
	l.Size = 0
}

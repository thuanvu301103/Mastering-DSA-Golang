package linked_list

// HasCycle checks if the linked list contains a loop (cycle).
// It returns true if a cycle is detected, false otherwise.
func (l *SinglyLinkedList) HasCycle() bool {
	if l.Head == nil || l.Head.Next == nil {
		return false
	}

	slow := l.Head
	fast := l.Head

	// Move fast by 2 steps and slow by 1 step
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// If they meet at the same node, there's a cycle
		if slow == fast {
			return true
		}
	}

	// If fast reaches nil, there is no cycle
	return false
}

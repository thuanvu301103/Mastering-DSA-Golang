package linked_list

func (l *SinglyLinkedList) FindMiddle() (*Node, bool) {
	if l.IsEmpty() {
		return nil, false
	}

	slow := l.Head
	fast := l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow, true
}

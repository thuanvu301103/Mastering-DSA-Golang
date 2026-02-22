package stack_queue

type CircularQueue struct {
	data  []int
	head  int
	tail  int
	size  int
	count int
}

func NewCircularQueue(k int) *CircularQueue {
	return &CircularQueue{
		data: make([]int, k),
		size: k,
		head: 0,
		tail: -1, // Tail starts at -1 for easier math
	}
}

func (q *CircularQueue) Enqueue(v int) bool {
	if q.IsFull() {
		return false
	}
	// Wrap around the tail index
	q.tail = (q.tail + 1) % q.size
	q.data[q.tail] = v
	q.count++
	return true
}

func (q *CircularQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	val := q.data[q.head]
	// Wrap around the head index
	q.head = (q.head + 1) % q.size
	q.count--
	return val, true
}

func (q *CircularQueue) IsFull() bool {
	return q.count == q.size
}

func (q *CircularQueue) IsEmpty() bool {
	return q.count == 0
}

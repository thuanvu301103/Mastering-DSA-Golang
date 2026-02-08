package sort

type MaxHeap struct {
	data []int
}

// Insert adds a value and bubbles it up to maintain heap property
func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.bubbleUp(len(h.data) - 1)
}

// ExtractMax removes and returns the largest element
func (h *MaxHeap) ExtractMax() int {
	if len(h.data) == 0 {
		panic("Heap is empty")
	}

	max := h.data[0]
	lastIndex := len(h.data) - 1

	// Move last element to root and bubble down
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]

	if len(h.data) > 0 {
		h.bubbleDown(0)
	}
	return max
}

func (h *MaxHeap) Peek() int {
	return h.data[0]
}

func (h *MaxHeap) GetArray() []int {
	return h.data
}

// Internal helpers
func (h *MaxHeap) bubbleUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[index] <= h.data[parent] {
			break
		}
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
	}
}

func (h *MaxHeap) bubbleDown(index int) {
	lastIndex := len(h.data) - 1
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left <= lastIndex && h.data[left] > h.data[largest] {
			largest = left
		}
		if right <= lastIndex && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
	}
}

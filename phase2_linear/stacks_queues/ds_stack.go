package stack_queue

type Stack struct {
	data []int
}

func (s *Stack) Push(n int) {
	s.data = append(s.data, n)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) <= 0 {
		return 0, false
	}

	index := len(s.data) - 1
	element := s.data[index]
	s.data = s.data[:index]

	return element, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.data) <= 0 {
		return 0, false
	}

	index := len(s.data) - 1
	return s.data[index], true

}

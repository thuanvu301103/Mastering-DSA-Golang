package stack_queue

type MonotonicStack struct {
	data []int
	Type string // "increasing" or "decreasing"
}

func (ms *MonotonicStack) Push(v int) {
	for len(ms.data) > 0 {
		top := ms.data[len(ms.data)-1]
		// Check for violation of the monotonic property
		if ms.Type == "increasing" && top > v {
			ms.data = ms.data[:len(ms.data)-1] // Pop
		} else if ms.Type == "decreasing" && top < v {
			ms.data = ms.data[:len(ms.data)-1] // Pop
		} else {
			break
		}
	}
	ms.data = append(ms.data, v)
}

func (ms *MonotonicStack) GetData() []int {
	return ms.data
}

func SolveDailyTemperatures(temps []int) []int {
	n := len(temps)
	result := make([]int, n)
	stack := []int{} // This will store INDICES

	for i := 0; i < n; i++ {
		// While current temp is warmer than the temp at the stack's top index
		for len(stack) > 0 && temps[i] > temps[stack[len(stack)-1]] {
			prevIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]  // Pop
			result[prevIdx] = i - prevIdx // Distance to the warmer day
		}
		stack = append(stack, i)
	}
	return result
}

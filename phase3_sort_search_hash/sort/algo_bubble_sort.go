package sort

func BubbleSort(array []int) []int {
	n := len(array)
	if n <= 1 {
		return array
	}

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return array
}

package sort

func SelectionSort(array []int) []int {
	n := len(array)
	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
	return array
}

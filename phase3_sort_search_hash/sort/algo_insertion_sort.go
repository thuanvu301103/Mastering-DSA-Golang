package sort

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1

		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = key
	}
	return array
}

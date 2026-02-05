package arrays

type DynamicArrayOps struct{}

func (o DynamicArrayOps) AppendElement(arr []int, item int) []int {
	var new_arr []int
	new_arr = append(arr, item)
	return new_arr
}

func (o DynamicArrayOps) FilterEvenNumbers(arr []int) []int {
	var result []int = []int{}
	for _, v := range arr {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func (o DynamicArrayOps) GetRange(arr []int, start, end int) []int {
	if end > len(arr) || start > end {
		return []int{}
	}

	return arr[start:end]
}

func (o DynamicArrayOps) ModifyElement(arr []int, index, value int) {
	arr[index] = value
}

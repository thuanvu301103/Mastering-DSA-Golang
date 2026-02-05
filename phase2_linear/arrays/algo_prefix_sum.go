package arrays

type PrefixSumOps struct{}

func (o PrefixSumOps) BuildPrefixArray(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return prefix
}

func (o PrefixSumOps) GetRangeSum(prefix []int, i, j int) int {
	return prefix[j+1] - prefix[i]
}

func (o PrefixSumOps) Build2DPrefixMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	p := make([][]int, rows+1)
	for i := range p {
		p[i] = make([]int, cols+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			p[r+1][c+1] = matrix[r][c] + p[r][c+1] + p[r+1][c] - p[r][c]
		}
	}
	return p
}

func (o PrefixSumOps) GetRectSum(p [][]int, r1, c1, r2, c2 int) int {
	return p[r2+1][c2+1] - p[r1][c2+1] - p[r2+1][c1] + p[r1][c1]
}

package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	col := len(nums)
	row := len(nums[0])
	if col*row != r*c {
		return nums
	}

	arr := make([]int, col*row)
	index := 0
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			arr[index] = nums[i][j]
			index++
		}
	}

	index = 0
	res := make([][]int, r*c)

	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
		for j := 0; j < c; j++ {
			res[i][j] = arr[index]
			index++
		}
	}
	return res
}

func main() {

}

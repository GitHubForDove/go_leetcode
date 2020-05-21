package main

import "fmt"

func countSquares(matrix [][]int) int {
	col := len(matrix)
	row := len(matrix[0])
	if matrix == nil || col == 0 || row == 0 {
		return 0
	}
	dp := make([][]int, len(matrix))
	fmt.Printf("col size= %d, row size= %d\n", col, row)
	// 二维数组初始化
	for i := 0; i < col; i++ {
		dp[i] = make([]int, row)
		for j := 0; j < row; j++ {
			dp[i][j] = matrix[i][j]
		}
	}

	result := 0
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 1 {
				dp[i][j] = Min(dp[i][j-1], Min(dp[i-1][j-1], dp[i-1][j])) + 1
			}
			result += dp[i][j]
		}
	}

	return result
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}
	res := countSquares(matrix)
	fmt.Println(res)
}

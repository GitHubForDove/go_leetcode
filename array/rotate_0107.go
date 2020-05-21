package main

import "fmt"

/**
面试题 01.07 旋转矩阵
*/

func rotate(matrix [][]int) {
	col := len(matrix)
	row := len(matrix[0])

	for i := 0; i < col; i++ {
		for j := i; j < row; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < col; i++ {
		for j := 0; j < row/2; j++ {
			matrix[i][j], matrix[i][row-j-1] = matrix[i][row-j-1], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

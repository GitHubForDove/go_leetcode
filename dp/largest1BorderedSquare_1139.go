package main

import "fmt"

/**
1139. 最大的以 1 为边界的正方形

给你一个由若干 0 和 1 组成的二维网格 grid，请你找出边界全部由 1 组成的最大 正方形 子网格，
并返回该子网格中的元素数量。如果不存在，则返回 0。


因为是找围成正方形的最大面积 所以只需要考虑 全部是1连起来的边长是多少
也就是说 只考虑 上面的是否为1 和 左边的是否为1
所以设置为 dp[i][j][k] --> k == 2
	k == 0 表示 i，j 左边的情况
	k == 1 表示 i，j 右边的情况
*/

func largest1BorderedSquare(grid [][]int) int {

	col := len(grid)
	row := len(grid[0])
	ans := 0
	dp := make([][][]int, col+1)
	for i := 0; i <= col; i++ {
		dp[i] = make([][]int, row+1)
		for j := 0; j <= row; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 1; i <= col; i++ {
		for j := 1; j <= row; j++ {
			// 如果是0则不要继续了
			if grid[i-1][j-1] == 0 {
				continue
			}
			dp[i][j][0] = dp[i-1][j][0] + 1
			dp[i][j][1] = dp[i][j-1][1] + 1
			minNum := min(dp[i][j][0], dp[i][j][1])
			for k := 0; k < minNum; k++ {
				// 无法确认  是否最小值 就是 目前边长的最大值 因为可能另外的两条边比较短 所以要进行比较
				if dp[i-k][j][1] >= k+1 && dp[i][j-k][0] >= k+1 {
					ans = max(ans, k+1)
				}
			}
		}
	}
	return ans * ans
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	res := largest1BorderedSquare(grid)
	fmt.Println(res)
}

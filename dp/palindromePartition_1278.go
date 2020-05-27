package main

import (
	"fmt"
	"math"
)

/**
1278. 分割回文串 III
给你一个由小写字母组成的字符串 s，和一个整数 k。

请你按下面的要求分割字符串：

首先，你可以将 s 中的部分字符修改为其他的小写英文字母。
接着，你需要把 s 分割成 k 个非空且不相交的子串，并且每个子串都是回文串。
请返回以这种方式分割字符串所需修改的最少字符数。

示例 1：

输入：s = "abc", k = 2
输出：1
解释：你可以把字符串分割成 "ab" 和 "c"，并修改 "ab" 中的 1 个字符，将它变成回文串。

示例 2：

输入：s = "aabbc", k = 3
输出：0
解释：你可以把字符串分割成 "aa"、"bb" 和 "c"，它们都是回文串。
示例 3：

输入：s = "leetcode", k = 8
输出：0
*/

/**
第一种解法 暴力搜索 + 记忆化结果
*/

var cache [][]int // cache记录 前i个字符组成的字符串 切割出j个 字符串所花费的最小代价
func palindromePartition(s string, k int) int {
	lens := len(s)
	if lens <= k {
		return 0
	}
	cache = make([][]int, lens)
	for i := 0; i < lens; i++ {
		cache[i] = make([]int, k)
		for j := 0; j < k; j++ {
			cache[i][j] = math.MinInt32
		}
	}
	cs := []byte(s)
	return dfs(cs, 0, k-1)

}

func dfs(cs []byte, start int, k int) int {

	// len - start代表当前要切割的字符串的长度
	// 如果这个长度小于等于要切割的次数意味着没办法继续切割了
	if len(cs)-start <= k {
		return -1
	}

	if cache[start][k] != math.MinInt32 {
		return cache[start][k]
	}

	// k <= 0 时 还有剩余的字符串  直接计算最小代价
	if k <= 0 {
		cache[start][k] = cost(cs, start, len(cs))
		return cache[start][k]
	}
	minCost := math.MaxInt32
	for i := 1; i <= len(cs); i++ {
		// 计算 start 到 start + i 的最小代价
		cost := cost(cs, start, start+i)
		// 接着计算后面字符串的最小代价
		sub := dfs(cs, start+i, k-1)
		if sub == -1 {
			break
		}
		minCost = min(minCost, cost+sub)
	}
	// 将最小代价记录下来
	cache[start][k] = minCost
	return cache[start][k]
}

func cost(cs []byte, start int, end int) int {
	res := 0
	left := start
	right := end - 1
	for left < right {
		if cs[left] != cs[right] {
			res++
		}
		left++
		right--
	}
	return res

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
使用动态规划解决问题
*/

func palindromePartition2(s string, k int) int {
	lens := len(s)
	if lens <= k {
		return 0
	}
	cs := []byte(s)
	dp := make([][]int, lens+1)
	// 计算i 到j 之间所需要的最小代价 并记录下来
	cost := make([][]int, lens)
	for i := 0; i < lens; i++ {
		cost[i] = make([]int, lens)
		for j := i + 1; j < lens; j++ {
			cost[i][j] = helper(cs, i, j)
		}
	}

	// 处理 k == 1的情况
	for i := 0; i <= lens; i++ {
		dp[i] = make([]int, k+1)
		if i > 0 {
			dp[i][1] = cost[0][i-1]
		}
	}

	for x := 2; x <= k; x++ {
		for i := x; i <= lens; i++ {
			dp[i][x] = lens + 1
			for j := x - 1; j < i; j++ {
				dp[i][x] = min(dp[i][x], dp[j][x-1]+cost[j][i-1])
			}
		}
	}
	return dp[lens][k]
}

func helper(cs []byte, left int, right int) int {
	res := 0
	for left < right {
		if cs[left] != cs[right] {
			res++
		}
		left++
		right--
	}
	return res
}

func main() {

	/*"aabbc"
	3*/
	s := "aabbc"
	k := 3
	res := palindromePartition2(s, k)
	fmt.Println(res)
}

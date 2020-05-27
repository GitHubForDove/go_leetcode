package main

import "fmt"

/**
双指针解决 最长回文子串
*/
func longestPalindrome(s string) string {
	lens := len(s)
	if lens <= 1 {
		return s
	}
	cs := []byte(s)
	var res string

	for i := 0; i < lens; i++ {
		s1 := palindrome(cs, i, i)
		s2 := palindrome(cs, i, i+1)
		if len(s1) > len(s2) && len(res) < len(s1) {
			res = s1
		} else if len(s1) < len(s2) && len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

/*func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}*/

func palindrome(cs []byte, left int, right int) string {
	//r := string(cs[left])
	for left >= 0 && right < len(cs) && cs[left] == cs[right] {
		left--
		right++
	}
	if left+1 < len(cs) && right-1 >= 0 && cs[left+1] == cs[right-1] {
		return string(cs[left+1 : right])
	} else {
		return ""
	}
}

/**
使用动态规划 解决
最长的回文子串 由小的回文子串构成 所以 存在重复子问题  可以使用动态规划的思想解决问题
基本思路：
dp[i][j] -->> i： 表示 起始字符串位置  j 表示： 字符串结束位置
dp[i][j] = dp[i+1][j-1]
*/

func longestPalindrome2(s string) string {
	lens := len(s)
	if lens <= 1 {
		return s
	}
	start := 0
	end := 0
	dp := make([][]bool, lens)
	for i := 0; i < lens; i++ {
		dp[i] = make([]bool, lens)
	}
	cs := []byte(s)
	for i := 1; i < lens; i++ {
		for j := 0; j < i; j++ {
			// (i-j) <= 2 是因为 cs[i]==cs[j] 且 cs[i] == cs[j] i-j <= 2 说明 i和j的位置 差距小于3 只有三种情况 a aa aba
			if cs[i] == cs[j] && ((i-j <= 2) || dp[j+1][i-1]) {
				dp[j][i] = true
				if (i - j) > (end - start) {
					start = j
					end = i
				}
			}
		}
	}
	return string(cs[start : end+1])
}

func main() {
	s := "aaaa"
	res := longestPalindrome2(s)
	fmt.Println(res)
}

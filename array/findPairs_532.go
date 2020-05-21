package main

import (
	"fmt"
	"sort"
)

/**
532. 数组中的K-diff数对

给定一个整数数组和一个整数 k, 你需要在数组里找到不同的 k-diff 数对。
这里将 k-diff 数对定义为一个整数对 (i, j), 其中 i 和 j 都是数组中的数字，且两数之差的绝对值是 k.

i - j == k  or  i - j = -k ->  j-i = k
*/
func findPairs(nums []int, k int) int {

	if nums == nil {
		return 0
	}

	if k < 0 {
		return 0
	}

	numsHas := make(map[int]bool)
	diffHas := make(map[int]bool)

	// 3, 1, 4, 1, 5
	// (1,3) (3,5)
	// 存较小的  用来去重
	for _, num := range nums {
		if numsHas[num-k] {
			diffHas[num-k] = true
		}
		if numsHas[num+k] {
			diffHas[num] = true
		}
		numsHas[num] = true
	}
	return len(diffHas)
}

func findPairs2(nums []int, k int) int {

	if k < 0 {
		return 0
	}

	sort.Ints(nums)
	lens := len(nums)
	res := 0
	prev := nums[0] - 1
	for i := 0; i < len(nums); i++ {
		if prev == nums[i] {
			continue
		}
		for j := i + 1; j < lens; j++ {
			if nums[j]-nums[i] > k {
				break
			}
			if nums[i]+k == nums[j] {
				res++
				break
			}
		}
		prev = nums[i]
	}
	return res
}

func main() {
	nums := []int{3, 1, 4, 1, 5,}
	k := 2
	res := findPairs2(nums, k)
	fmt.Println(res)
}

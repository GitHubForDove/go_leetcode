package main

import (
	"fmt"
)

func arrayNesting(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	visited := make(map[int]bool, lens)
	result := 0

	for i := 0; i < lens; i++ {
		// 不需要重复计算  访问过的 已经不可能存在比现在更大的可能
		if !visited[i] {
			index := nums[i]
			flag := 0
			count := 0
			for {
				if nums[i] == index {
					flag++
					if flag == 2 {
						break
					}
				}

				index = nums[index]
				count++
				visited[index] = true
			}

			result = max(count, result)
		}
	}
	return result
}

/**
	并查集解决
*/
func arrayNesting2(nums []int) int {
	lens := len(nums)
	fat = make([]int, lens)
	// 初始化并查集
	for i := 0; i < lens; i++ {
		fat[i] = i
	}

	for i := 0; i < lens; i++ {
		// fat[find(i)] : i 下标的祖先对应得下标  find(nums[i])  nums[i] 对应得祖先
		fat[find(i)] = find(nums[i])
	}

	res := 0

	m := make(map[int]int)
	for i := 0; i < lens; i++ {
		if _, ok := m[find(i)]; !ok {
			m[find(i)] = 1
		} else {
			m[find(i)]++
		}
		res = max(res, m[find(i)])
	}
	return res
}

var fat []int

func find(x int) int {
	if x == fat[x] {
		return x
	}
	fat[x] = find(fat[x])
	return fat[x]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{
		//5, 4, 0, 3, 1, 6, 2,
		0, 1, 2,
	}

	res := arrayNesting2(nums)
	fmt.Println(res)
}

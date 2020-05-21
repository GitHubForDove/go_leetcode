package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) <= 0 {
		return res
	}
	// 排序
	sort.Ints(nums)
	lens := len(nums)
	for i := 0; i < lens; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := lens - 1
		for left < right {
			// 去重
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			tmpSum := nums[i] + nums[left] + nums[right]
			if tmpSum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if tmpSum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func main() {
	nums := []int{
		1, -1, -1, 0,
	}
	res := threeSum(nums)
	fmt.Println(res)
}

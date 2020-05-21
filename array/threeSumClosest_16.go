package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	lens := len(nums)
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < lens; i++ {
		left := i + 1
		right := lens - 1
		for left < right {
			tmpSum := nums[i] + nums[left] + nums[right]
			if Abs(tmpSum-target) < Abs(target-ans) {
				ans = tmpSum
			}
			if target-tmpSum < 0 {
				right--
			} else if target-tmpSum > 0 {
				left++
			} else {
				return ans
			}
		}
	}

	return ans
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	nums := []int{
		0, 0, 0,
	}

	res := threeSumClosest(nums, 1)
	fmt.Println(res)
}

package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	lens := len(nums)
	if lens < 4 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < lens; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < lens; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := lens - 1

			for left < right {
				tmpSum := nums[i] + nums[j] + nums[left] + nums[right]
				if tmpSum < target {
					left++
				} else if tmpSum > target {
					right--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{
		-1, 2, 2, -5, 0, -1, 4,
	}
	res := fourSum(nums, 3)
	fmt.Println(res)
}

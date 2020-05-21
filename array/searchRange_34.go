package main

import "fmt"

func searchRange(nums []int, target int) []int {
	lens := len(nums)
	if lens == 0 {
		return []int{-1, -1}
	}

	left := 0
	right := lens
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left >= lens && nums[left] != target {
		return []int{-1, -1}
	}
	cur := left
	var index1 int
	var index2 int
	for cur >= 0 && nums[cur] == target {
		index1 = cur
		cur--
	}
	cur = left
	for cur < lens && nums[cur] == target {
		index2 = cur
		cur++
	}
	return []int{index1, index2}
}

func main() {
	nums := []int{
		1,
	}
	target := 1
	result := searchRange(nums, target)
	fmt.Println(result)
}

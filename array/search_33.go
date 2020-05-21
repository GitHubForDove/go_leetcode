package main

import "fmt"

func search(nums []int, target int) int {
	lens := len(nums)
	if lens == 0 {
		return -1
	}

	left := 0
	right := lens - 1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < nums[right] {
			if nums[mid+1] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			// 右边有序
			if nums[mid] >= target && target >= nums[left] {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func main() {
	nums := []int{
		4, 5, 6, 7, 0, 1, 2,
	}
	target := 0
	index := search(nums, target)
	fmt.Println(index)
}

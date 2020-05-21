package main

import "fmt"

func removeDuplicates(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}

	left := 0
	right := 1
	for right < lens {
		if nums[right] == nums[right-1] {
			right++
		} else {
			left++
			nums[left] = nums[right]
			right++
		}
	}

	return left + 1
}

func main() {
	nums := []int{
		1,1,2,
		//0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
	}
	l := removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(l)
}

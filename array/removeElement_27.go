package main

import "fmt"

func removeElement(nums []int, val int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}

	left := 0
	right := 0

	for right < lens {
		if nums[right] == val {
			right++
		} else {
			nums[left] = nums[right]
			left++
			right++
		}
	}
	return left

}

func main() {
	nums := []int{
		//3,2,2,3,
		0, 1, 2, 2, 3, 0, 4, 2,
	}
	val := 2
	l := removeElement(nums, val)
	fmt.Println(nums)
	fmt.Println(l)
}

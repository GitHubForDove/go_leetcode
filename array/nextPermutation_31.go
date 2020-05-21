package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	lens := len(nums)
	if lens == 0 {
		return
	}

	left := lens - 2
	for left >= 0 {
		if nums[left] >= nums[left+1] {
			left--
		} else {
			minIndex := left + 1
			for i := minIndex; i < lens; i++ {
				if nums[i] > nums[left] && nums[minIndex] > nums[i] {
					minIndex = i
				}
			}
			nums[left], nums[minIndex] = nums[minIndex], nums[left]
			tmp := nums[left+1:]
			sort.Ints(tmp)
			break
		}
	}
	if left == -1 {
		sort.Ints(nums)
	}
}


func nextPermutation2(nums []int) {
	lens := len(nums)
	if lens == 0 {
		return
	}

	
}

func main() {
	nums := []int{
		//3, 2, 1,
		5, 1, 1,
	}
	nextPermutation(nums)
	fmt.Println(nums)
}

package main

import "fmt"

func twoSum(nums []int, target int) []int {

	result := make([]int, 2)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[target - nums[i]]
		if ok  {
			result[0] = m[target - nums[i]]
			result[1] = i
			break
		}
		m[nums[i]] = i
	}
	return result
}

func main() {
	nums := []int{
		2, 7, 11, 15,
	}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}

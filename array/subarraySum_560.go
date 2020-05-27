package main

import "fmt"

/**
前缀和
*/
func subarraySum(nums []int, k int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}

	prefix := make([]int, lens+1)
	sum := 0
	for i := 0; i < lens; i++ {
		sum += nums[i]
		prefix[i+1] = sum
	}
	right := len(prefix) - 1
	count := 0
	for i := 0; i < len(prefix); i++ {

		for i < right {
			tmpSum := prefix[right] - prefix[i]
			if tmpSum == k {
				count++
			}
			right--
		}
		right = len(prefix) - 1
	}
	return count
}

/**
前缀和 加 map
*/
func subarraySum2(nums []int, k int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}

	sum := 0
	var m map[int]int
	m = make(map[int]int)
	m[0] = 1
	prefixSum := make([]int, lens+1)
	count := 0

	for i := 0; i < lens; i++ {
		sum += nums[i]
		prefixSum[i+1] = sum
		if _, ok := m[sum - k]; ok {
			count += m[sum - k]
		}

		if _, ok := m[sum]; ok {
			m[sum]++
		} else {
			m[sum] = 1
		}
	}

	return count
}

func main() {
	nums := []int{
		-1, -1, 1,
	}
	k := 0

	res := subarraySum2(nums, k)
	fmt.Println(res)
}

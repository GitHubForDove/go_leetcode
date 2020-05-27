package main

import "fmt"

/**
	前缀和
*/
func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	hashmap := make(map[int]int)
	hashmap[0] = -1 //对0特殊处理
	sum := 0
	for i := 0;i < l;i++ {
		sum += nums[i]
		if k != 0 {
			sum = sum % k
		}

		if v,ok := hashmap[sum]; ok {
			if i - v > 1 { //必须有2个数
				return true
			}
		} else {
			hashmap[sum] = i
		}
	}
	return false

}

func main() {
	nums := []int{
		//23, 2, 4, 6, 7,
		//0, 0,
		5, 0, 0,
	}
	k := 0
	res := checkSubarraySum(nums, k)
	fmt.Println(res)
}

package main

import "fmt"

/**
	面试题：08.04. 幂集
	输入： nums = [1,2,3]
	幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。

	说明：解集不能包含重复的子集

 输出：
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

func subsets(nums []int) [][]int {
	lens := len(nums)
	if lens == 0 {
		return nil
	}
	//var result [][]int
	var list []int
	result = make([][]int, 0)
	list = make([]int, 0)
	dfs(nums, list, 0)
	return result
}

var result [][]int

func dfs(nums []int, list []int, index int) {
	tmp := make([]int, len(list))
	copy(tmp, list)
	result = append(result, tmp)

	for i := index; i < len(nums); i++ {
		list = append(list, nums[i])
		dfs(nums, list, i+1)
		list = list[:len(list)-1]
	}

}

/**
使用位运算
*/
func subsets2(nums []int) [][]int {
	lens := len(nums)
	if lens == 0 {
		return nil
	}
	n := 1 << lens
	var result [][]int
	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < lens; j++ {
			if ((i >> j) & 1) == 1 {
				tmp = append(tmp, nums[j])
			}
		}
		result = append(result, tmp)
	}
	return result
}

func main() {
	nums := []int{
		1, 2, 3,
	}
	result := subsets2(nums)
	fmt.Println(result)
}

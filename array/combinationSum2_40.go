package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	lens := len(candidates)
	sort.Ints(candidates)
	result := make([][]int, 0)
	if lens == 0 {
		return result
	}
	list := make([]int, 0)
	dfs(candidates, target, list, 0, 0, &result)
	return result
}

func dfs(candidates []int, target int, list [] int, sum int, index int, result *[][]int) {
	if sum > target {
		return
	}
	if sum == target {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
		return
	}

	for i := index; i < len(candidates); i++ {

		//剪枝
		if target < candidates[i] {
			continue
		}

		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		list = append(list, candidates[i])
		dfs(candidates, target, list, sum+candidates[i], i+1, result)
		list = list[:len(list)-1]
	}
}

func main() {
	candidates := []int{
		10, 1, 2, 7, 6, 1, 5,
	}
	target := 8
	res := combinationSum2(candidates, target)
	fmt.Println(res)
}

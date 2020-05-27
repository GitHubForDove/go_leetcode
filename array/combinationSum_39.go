package main

import "fmt"

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	lens := len(candidates)
	if lens == 0 {
		return result
	}
	var list []int
	dfs(candidates, list, target, 0, 0)
	return result
}

func dfs(candidates []int, list []int, target int, sum int, index int) {
	if sum == target {
		tmp := make([]int, len(list))
		copy(tmp, list)
		result = append(result, tmp)
		return
	}

	for i := index; i < len(candidates); i++ {
		if sum+candidates[i] <= target {
			list = append(list, candidates[i])
			dfs(candidates, list, target, sum + candidates[i], i)
			list = list[:len(list)-1]
		}
	}
}

func main() {
	candidates := []int {
		2,3,5,
	}
	target := 8
	res := combinationSum(candidates, target)
	fmt.Println(res)
}

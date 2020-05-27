package main

import "fmt"

/**
	1011. 在 D 天内送达包裹的能力

	思路：
	使用二分查找，其实 找 能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。
	就是在 (max(weight[i] ~  sum(Weights[i])) 之间找
    最大的情况 就是 每批货 都按最大的货物重量来计算  就是每天只能运一批货
	最小的情况  就是 船很大  一次把货全部都运完了

*/

func shipWithinDays(weights []int, D int) int {
	lens := len(weights)
	sum := 0
	maxWeight := weights[0]
	for i := 0; i < lens; i++ {
		sum += weights[i]
		maxWeight = max(maxWeight, weights[i])
	}
	left := maxWeight
	right := sum

	for left < right {
		mid := (left + right) >> 1
		tmpSum := 0
		days := 1
		for _, w := range weights {
			tmpSum += w
			if tmpSum > mid {
				days++
				tmpSum = w
			}
		}
		if days > D {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weights := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	D := 5
	res := shipWithinDays(weights, D)
	fmt.Println(res)
}

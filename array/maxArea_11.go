package main

import "fmt"

func maxArea(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	max := 0

	for index, v := range height {
		left := v
		cur := index + 1
		for cur < len(height) {
			right := height[cur]
			area := Min(right, left) * (cur - index)
			max = Max(area, max)
			cur++
		}
	}

	return max
}

// Area = Max(min(height[i], height[j]) * (j-i)) {其中0 <= i < j < height,size()}

// 使用动态规划

func maxArea2(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		h := Min(height[i], height[j])
		max = Max(max, h * (j - i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{
		1, 8, 6, 2, 5, 4, 8, 3, 7,
	}
	res := maxArea(height)
	fmt.Println(res)
}

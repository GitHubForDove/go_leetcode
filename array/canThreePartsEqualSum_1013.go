package main

import (
	"fmt"
)

/**
leetcode 1013 将数组划分为三个相等的部分

给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。

形式上，如果可以找出索引 i+1 < j 且满足 
(A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 
就可以将数组三等分。

*/

func canThreePartsEqualSum(A []int) bool {
	if A == nil {
		return true
	}

	var sum int
	for _, v := range A {
		sum += v
	}

	if sum%3 != 0 {
		return false
	}

	avg := sum / 3
	var n int
	var count int
	for i := 0; i < len(A); i++ {
		n += A[i]
		sum -= A[i]
		if n == avg {
			n = 0
			count++
		}
	}
	return sum == 0 && count >= 3
}

func main() {
	A := []int{
		//0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1,
		//0,2,1,-6,6,7,9,-1,2,0,1,
		//3,3,6,5,-2,2,5,1,-9,4,
		//1, -1, 1, -1,
		10, -10, 10, -10, 10, -10, 10, -10,
	}
	b := canThreePartsEqualSum(A)
	fmt.Println(b)
}

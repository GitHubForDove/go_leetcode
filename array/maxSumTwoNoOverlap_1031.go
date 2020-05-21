package main

import "fmt"


/**
	1031. 两个非重叠子数组的最大和

	给出非负整数数组 A ，返回两个非重叠（连续）子数组中元素的最大和，子数组的长度分别为 L 和 M。
	（这里需要澄清的是，长为 L 的子数组可以出现在长为 M 的子数组之前或之后。）

	从形式上看，返回最大的 V，而 V = (A[i] + A[i+1] + ... + A[i+L-1]) + (A[j] + A[j+1] + ... + A[j+M-1])
	并满足下列条件之一：
		0 <= i < i + L - 1 < j < j + M - 1 < A.length, 或
		0 <= j < j + M - 1 < i < i + L - 1 < A.length.


	sum := 0
	for i := 0; i < lens; i++ {
		sum += A[i]
		prefixSum[i+1] = sum
	}

	前缀和  ： 注意 前缀和的下标指向的数值
 */
func maxSumTwoNoOverlap(A []int, L int, M int) int {

	lens := len(A)
	// 前缀和
	prefixSum := make([]int, lens+1)
	sum := 0
	for i := 0; i < lens; i++ {
		sum += A[i]
		prefixSum[i+1] = sum
	}
	result := 0
	for i := L; i < len(prefixSum); i++ {
		// 长度为L 的和
		LSum := prefixSum[i] - prefixSum[i-L]
		MSum := 0
		// 如果M窗口在左边
		for j := M; j <= i-L; j++ {
			MSum = max(MSum, prefixSum[j]-prefixSum[j-M])
		}

		// 如果M窗口在右边
		for j := len(prefixSum) - 1; j >= i+M; j-- {
			MSum = max(MSum, prefixSum[j]-prefixSum[j-M])
		}
		result = max(result, MSum+LSum)
	}
	return result
}


func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	A := []int{
		//0, 6, 5, 2, 2, 5, 1, 9, 4,
		2, 1, 5, 6, 0, 9, 5, 0, 3, 8,
	}
	L := 4
	M := 3
	result := maxSumTwoNoOverlap(A, L, M)
	fmt.Println(result)
}

package main

/**
974. 和可被 K 整除的子数组

思路：前缀和 超时
*/

func subarraysDivByK(A []int, K int) int {
	lens := len(A)
	if lens == 0 {
		return 0
	}

	prefixSum := make([]int, len(A)+1)
	sum := 0
	for i := 1; i <= len(A); i++ {
		sum += A[i-1]
		prefixSum[i] = sum

	}
	count := 0
	for i := 0; i < len(prefixSum); i++ {
		tmpSum := prefixSum[i]
		right := len(prefixSum)
		for i < right {
			if (prefixSum[right]-tmpSum)%K == 0 {
				count++
			}
			right--
		}
		right = len(prefixSum)
	}
	return count
}

/**
	有 A[:i] A[:j] 且 i < j， A[:i] % K == A[:j] % K
	那么就一定存在 A[i+1:j]%K == 0
 */
func subarraysDivByK2(A []int, K int) int {
	lens := len(A)
	if lens == 0 {
		return 0
	}

	modK := make([]int, K)
	prefixSum := 0
	modK[0] = 1
	count := 0

	for _, v := range A {
		prefixSum = (prefixSum + v) % K
		if prefixSum < 0 {
			prefixSum += K
		}
		count += modK[prefixSum]
		modK[prefixSum] ++
	}
	return count
}

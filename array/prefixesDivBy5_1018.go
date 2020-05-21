package main

func prefixesDivBy5(A []int) []bool {
	var result []bool
	lens := len(A)
	var n int

	for i := 0; i < lens; i++ {
		n = n * 2 + A[i]
		// 除以5  所有整个值指与 最后一位数  0或者 5 相关 所以直接余5 将数值缩小 防止越界
		n %= 5
		if n % 5 == 0 {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}

func main() {

}

package main

func maxScoreSightseeingPair(A []int) int {

	var max int
	var left int
	for i := 0; i < len(A); i++ {
		sum := A[i] - i
		if sum > max {
			max = sum
		}

		t := A[i] + i
		if left < i {
			left = t
		}

		if max == 0 {
			return 0
		}
 	}
	return max
}

func main() {

}

package main

func searchInsert(nums []int, target int) int {

	lens := len(nums)
	if lens == 0 {
		return 0
	}

	left := 0
	right := lens
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {

}

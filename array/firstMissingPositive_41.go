package main

func firstMissingPositive(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return -1
	}

	for i := 0; i < len(nums); i++ {
		for nums[i] >= 1 && nums[i] <= lens && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < lens; i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}
	return lens + 1
}

func main() {

}

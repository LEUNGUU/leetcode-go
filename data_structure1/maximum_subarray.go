package data_structure1

func maxSubArray(nums []int) int {
	max_sum = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		max_sum = max(nums[i], max_sum)
	}
	return max_sum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

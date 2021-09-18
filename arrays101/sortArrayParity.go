package arrays101

func sortArrayByParity(nums []int) []int {
	s := 0
	for c := 0; c < len(nums); c++ {
		if nums[c]%2 == 0 {
			nums[s], nums[c] = nums[c], nums[s]
			s++
		}
	}
	return nums
}

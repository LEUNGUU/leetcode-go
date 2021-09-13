package arrays101

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 0
	for {
		if i <= j && j < len(nums) {
			if nums[i] != nums[j] {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		} else {
			i++
			break
		}
	}
	return i
}

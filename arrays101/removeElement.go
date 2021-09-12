package arrays101

func removeElement(nums []int, val int) int {
	times := 0
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			if nums[j] == val {
				j--
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				j--
				i++
				times++
			}
		} else {
			i++
			times++
		}
	}
	return times
}

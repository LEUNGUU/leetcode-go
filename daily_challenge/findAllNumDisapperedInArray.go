package daily_challenge

func myAbs(a int) int {
	if a > 0 {
		return a
	} else {
		return a * -1
	}
}

func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[myAbs(nums[i])-1] > 0 {
			nums[myAbs(nums[i])-1] *= -1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

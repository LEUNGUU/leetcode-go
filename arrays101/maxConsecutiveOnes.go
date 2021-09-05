package arrays101

func findMaxConsecutiveOnes(nums []int) int {
	res, count := 0, 0
	for _, val := range nums {
		if val == 1 {
			count++
		} else {
			res = bigger(res, count)
			count = 0
		}
	}
	if count != 0 {
		res = bigger(res, count)
	}
	return res
}

func bigger(res, count int) int {
	if res > count {
		return res
	} else {
		return count
	}
}

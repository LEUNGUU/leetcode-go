package arrays101

func findNumbers(nums []int) int {
	res := 0
	for _, val := range nums {
		count := 0
		for 1 > 0 {
			if val/10 != 0 {
				count++
				val = val / 10
			} else {
				break
			}
		}
		if count%2 > 0 {
			res++
		}
	}
	return res
}

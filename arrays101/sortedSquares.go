package arrays101

import "math"

func sortedSquares(nums []int) []int {
	var res []int
	for left, right := 0, len(nums)-1; left <= right; {
		if math.Abs(float64(nums[left])) >= math.Abs(float64(nums[right])) {
			res = append(res, nums[left]*nums[left])
			left += 1
		} else {
			res = append(res, nums[right]*nums[right])
			right -= 1
		}
	}
	length := len(res) - 1
	for idx := 0; idx <= length/2; idx++ {
		res[idx], res[length-idx] = res[length-idx], res[idx]
	}
	return res
}

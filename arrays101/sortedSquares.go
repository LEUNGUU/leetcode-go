package arrays101

func sortedSquares(nums []int) []int {
	var square_nums []int
	for _, num := range nums {
		square_nums = append(square_nums, num*num)
	}
	for i := len(square_nums) - 1; i >= 0; i-- {
		if square_nums[i] < square_nums[0] {
			square_nums[i], square_nums[0] = square_nums[0], square_nums[i]
		}
	}
	return square_nums
}

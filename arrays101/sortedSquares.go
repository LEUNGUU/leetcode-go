package arrays101

func sortedSquares(nums []int) []int {
	var square_nums []int
	for _, num := range nums {
		square_nums = append(square_nums, num*num)
	}
	for i, j := 0, len(square_nums)-1; j == i; {
		if square_nums[j] < square_nums[i] {
			square_nums[i], square_nums[j] = square_nums[j], square_nums[i]
			j -= 1
		} else {
			j -= 1
		}
	}
	return square_nums
}

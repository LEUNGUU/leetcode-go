package arrays101

func moveZeroes(nums []int) {
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] == 0 {
	// 		for j := i + 1; j < len(nums); j++ {
	// 			if nums[j] != 0 {
	// 				nums[i], nums[j] = nums[j], nums[i]
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	// keep nonZero element on s's left size
	// keep zero element between s and c
	for s, c := 0, 0; c < len(nums); c++ {
		if nums[c] != 0 {
			nums[s], nums[c] = nums[c], nums[s]
			s++
		}
	}
}

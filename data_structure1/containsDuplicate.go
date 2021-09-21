package data_structure1

// func containsDuplicate(nums []int) bool {
// 	quickSort(nums, 0, len(nums)-1)
// 	for i:=1; i < len(nums); i++ {
// 		if nums[i] == nums[i-1] {
// 			return true
// 		}
// 	}
// 	return false
// }
//
//
// func quickSort(nums []int, left, right int) {
// 	if left < right {
// 	  nums, pi := partition(nums, left, right)
// 	  quickSort(nums, left, pi-1)
// 	  quickSort(nums, pi+1, right)
// 	}
// }
//
// func partition(nums []int, left, right int) ([]int, int) {
// 	pivot := nums[right]
// 	smaller := left
// 	for current:=left; current<len(nums); current++ {
// 		if nums[current] < pivot {
// 			nums[smaller], nums[current] = nums[current], nums[smaller]
// 			smaller++
// 		}
// 	}
// 	nums[smaller], nums[right] = nums[right], nums[smaller]
// 	return nums, smaller
// }

func containsDuplicate(nums []int) bool {
	golang_set := make(map[int]bool)
	for _, val := range nums {
		_, ok := golang_set[val]
		if ok {
			return true
		}
		golang_set[val] = true
	}
	return false
}

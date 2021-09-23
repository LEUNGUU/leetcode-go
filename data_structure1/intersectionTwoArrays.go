package data_structure1

func intersect(nums1 []int, nums2 []int) []int {
}

func quickSort(nums []int, left, right int) {
	num, pi := partition(nums, 0, len(nums)-1)
	quickSort(num, 0, pi-1)
	quickSort(num, pi+1, len(nums)-1)
}

func partition(nums []int, left, right int) ([]int, int) {
	if left < right {
		pivot := nums[right]

	}

}

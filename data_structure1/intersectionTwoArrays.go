package data_structure1

func intersect(nums1 []int, nums2 []int) []int {
	quickSort(nums1, 0, len(nums1)-1)
	quickSort(nums2, 0, len(nums2)-1)
	i, j, k := 0, 0, 0
	for {
		if i < len(nums1) && j < len(nums2) {
			switch {
			case nums1[i] < nums2[j]:
				i++
			case nums1[i] > nums2[j]:
				j++
			default:
				nums1[k] = nums1[i]
				k++
				i++
				j++
			}
		} else {
			break
		}
	}
	return nums1[:k]
}

func quickSort(nums []int, left, right int) {
	if left < right {
		nums, pi := partition(nums, left, right)
		quickSort(nums, left, pi-1)
		quickSort(nums, pi+1, right)
	}
}

func partition(nums []int, left, right int) ([]int, int) {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return nums, i
}

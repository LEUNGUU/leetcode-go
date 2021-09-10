package arrays101

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, p := m-1, n-1, m+n-1
	for {
		if i >= 0 && j >= 0 {
			if nums1[i] <= nums2[j] {
				nums1[p] = nums2[j]
				j--
				p--
			} else {
				nums1[p], nums1[i] = nums1[i], 0
				i--
				p--
			}
		} else {
			break
		}
	}
	if j >= 0 {
		for idx := 0; idx <= j; idx++ {
			nums1[idx] = nums2[idx]
		}
	}
}

package arrays101

func merge(nums1 []int, m int, nums2 []int, n int) {
	var res []int
	i, j := 0, 0
	for {
		if i <= m && j <= n {
			if nums1[i] <= nums2[j] {
				res = append(res, nums1[i])
				i++
			} else {
				res = append(res, nums2[j])
				j++
			}
		} else {
			break
		}
	}
	if i < m {
		res = append(res, nums1[i:]...)
	}
	if j < n {
		res = append(res, nums2[j:]...)
	}
	nums1 = res
}

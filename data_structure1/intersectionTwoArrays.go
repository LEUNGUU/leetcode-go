package data_structure1

func intersect(nums1 []int, nums2 []int) []int {
	rc := []int{}
	var loopArr []int
	var another []int
	if len(nums1) > len(nums2) {
		loopArr := nums2
		another := nums1
	} else {
		loopArr := nums1
		another := nums2
	}
	for _, val := range loopArr {
		if containElem(another, val) {
			rc = append(rc, val)
		}
	}
	return rc
}

func containElem(arr []int, elem int) bool {
	for _, val := range arr {
		if val == elem {
			return true
		}
	}
	return false
}

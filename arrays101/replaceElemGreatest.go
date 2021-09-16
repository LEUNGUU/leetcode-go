package arrays101

func replaceElements(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], max = max, large(max, arr[i])
	}
	return arr
}

func large(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

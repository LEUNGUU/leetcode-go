package arrays101

func heightChecker(heights []int) int {
	copyHeight := make([]int, len(heights))
	copy(copyHeight, heights)
	quickSort(heights, 0, len(heights)-1)
	times := 0
	for idx, val := range heights {
		if copyHeight[idx] != val {
			times++
		}
	}
	return times
}

func quickSort(arr []int, left, right int) {
	if left < right {
		arr, p := partition(arr, left, right)
		quickSort(arr, left, p-1)
		quickSort(arr, p+1, right)
	}
}

func partition(arr []int, left, right int) ([]int, int) {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}

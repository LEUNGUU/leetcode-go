package arrays101

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	i := 0
	// left up
	for i < len(arr)-1 {
		if arr[i] < arr[i+1] {
			i++
		} else {
			break
		}
	}

	if i == 0 || i == len(arr)-1 {
		return false
	}

	// right down
	for i < len(arr)-1 {
		if arr[i] > arr[i+1] {
			i++
		} else {
			break
		}
	}
	return i == len(arr)-1
}

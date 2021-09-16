package arrays101

func duplicateZeros(arr []int) {
	// Notice: i := 0 only executes one time while initialization
	// the increment of i will be executed before the condition in every loop
	for idx := 0; idx < len(arr); idx++ {
		if arr[idx] == 0 {
			// discard the last element
			arr = arr[:len(arr)-1]
			// insert a zero
			var tmp = append(arr[:idx], 0)
			// concat them
			arr = append(tmp, arr[idx:]...)
			// skip the inserted one
			idx++
		}
	}
}

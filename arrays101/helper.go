package arrays101

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, val := range a {
		if val != b[idx] {
			return false
		}
	}
	return true
}

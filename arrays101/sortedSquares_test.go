package arrays101

import "testing"

func TestSortedSquares(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}
	want := []int{4, 9, 9, 49, 121}
	res := sortedSquares(nums)
	if !Equal(res, want) {
		t.Fatalf("testing sorted squares with %v, want: %v, but return: %v", nums, want, res)
	}
}

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

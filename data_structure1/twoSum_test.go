package data_structure1

import "testing"

func TestTwoSum(t *testing.T) {
	wants := []int{0, 3}
	tc := []int{2, 8, 11, 7}
	res := twoSum(tc, 9)
	if !Equal(res, wants) {
		t.Fatalf("testing twoSum, want: %v, but return: %v", wants, res)
	}
}

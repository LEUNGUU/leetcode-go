package arrays101

import (
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	tc := []int{1, 3, 2, 4}
	want := []int{2, 4, 1, 3}
	res := sortArrayByParity(tc)
	if !Equal(want, res) {
		t.Fatalf("sortArrayByParity error, want: %v, but return: %v", want, res)
	}
}

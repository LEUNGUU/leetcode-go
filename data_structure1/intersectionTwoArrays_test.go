package data_structure1

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	tc1 := []int{1, 2, 2, 1}
	tc2 := []int{2}
	wants := []int{2}
	res := intersect(tc1, tc2)
	if !Equal(res, wants) {
		t.Fatalf("test intersect, want: %v, but return: %v", wants, res)
	}
}

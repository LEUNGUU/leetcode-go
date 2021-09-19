package arrays101

import (
	"testing"
)

func TestHeightChecker(t *testing.T) {
	tc := []int{1, 3, 2, 4, 5}
	wants := 2
	res := heightChecker(tc)
	if res != wants {
		t.Fatalf("Test height checker, want: %v, but return: %v", wants, res)
	}
}

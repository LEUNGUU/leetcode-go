package arrays101

import (
	"testing"
)

func TestReplaceElements(t *testing.T) {
	var testcases []map[string][]int
	tc1 := map[string][]int{
		"arr":  {1, 4, 1, 76, 34, 2, 4},
		"want": {76, 76, 76, 34, 4, 4, -1},
	}
	tc2 := map[string][]int{
		"arr":  {1},
		"want": {-1},
	}
	testcases = append(testcases, tc1)
	testcases = append(testcases, tc2)
	for _, tc := range testcases {
		res := replaceElements(tc["arr"])
		if !Equal(res, tc["want"]) {
			t.Fatalf("testing replace elements with %v, want: %v, but return: %v", tc["arr"], tc["want"], res)
		}
	}

}

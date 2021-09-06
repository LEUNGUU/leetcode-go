package arrays101

import (
	"strconv"
	"strings"
	"testing"
)

var tc_set = map[string][]int{
	"tc1_3": {12, 345, 11, 1212},
	"tc2_1": {12, 111, 222},
	"tc3_1": {1212},
	"tc4_0": {2121211},
}

func TestEvenDigit(t *testing.T) {
	for k, v := range tc_set {
		want, _ := strconv.Atoi(strings.Split(k, "_")[1])
		res := findNumbers(v)
		if res != want {
			t.Fatalf("testing even Digit with %v, want: %v, but return: %v", v, want, res)
		}
	}
}

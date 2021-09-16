package arrays101

import (
	"strconv"
	"strings"
	"testing"
)

func TestValidMountainArray(t *testing.T) {
	var tcs = map[string][]int{
		"tc1_0": {12, 345, 11, 1212},
		"tc2_1": {0, 1, 0},
		"tc3_0": {1212},
		"tc4_0": {1, 2, 3, 4, 5},
	}
	for k, v := range tcs {
		want, _ := strconv.ParseBool(strings.Split(k, "_")[1])
		res := validMountainArray(v)
		if res != want {
			t.Fatalf("testing valid mountain array with %v, want: %v, but return: %v", v, want, res)
		}
	}
}

package data_structure1

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	want := true
	tc := "aab"
	tc1 := "aab"
	res := canConstruct(tc, tc1)
	if res != want {
		t.Fatalf("test canConstruct error, want: %v, but return: %v", want, res)
	}
}

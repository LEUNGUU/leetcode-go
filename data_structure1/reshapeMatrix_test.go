package data_structure1

import (
	"testing"
)

// TODO: Lazy to fix it.
func TestReshapeMatrix(t *testing.T) {
	wants := 2
	tc := [][]int{{1, 2, 3, 4}}
	res := matrixReshape(tc, 2, 2)
	if res != nil {
		t.Fatalf("Test reshape matrix error, want: %v, but return: %v", wants, res)
	}
}

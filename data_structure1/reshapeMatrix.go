package data_structure1

func matrixReshape(mat [][]int, r int, c int) [][]int {
	sz := len(mat) * len(mat[0])
	wanted := r * c
	if sz != wanted {
		return mat
	} else {
		all := []int{}
		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[0]); j++ {
				all = append(all, mat[i][j])
			}
		}
		res := make([][]int, r)
		for i := 0; i < r; i++ {
			res[i] = make([]int, c)
		}
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				res[i][j] = all[0]
				all = all[1:]
			}
		}
		return res
	}
}

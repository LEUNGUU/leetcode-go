package data_structure1

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	triangle[0] = make([]int, 1)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		triangle[i] = row
	}
	return triangle
}

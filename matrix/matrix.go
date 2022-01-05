package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true iff a and b are adjacent in lst.
func AreAdjacent(lst []int, a, b int) bool {
	for i := 0; i < len(lst)-1; i++ {
		if (lst[i] == a && lst[i+1] == b) || (lst[i] == b && lst[i+1] == a) {
			return true
		}
	}
	return false
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	// Check if mat is nil.
	if mat == nil || len(mat) == 0 {
		return mat
	}
	transposed := [][]int{}
	for i := 0; i < len(mat[0]); i++ {
		var row []int
		for j := 0; j < len(mat); j++ {
			row = append(row, mat[j][i])
		}
		transposed = append(transposed, row)
	}
	return transposed
}

// AreNeighbors returns true iff a and b are neighbors in the 2D matrix mat.
func AreNeighbors(mat [][]int, a, b int) bool {
	// Check if mat is nil.
	if mat == nil || len(mat) == 0 {
		return false
	}
	// Check if a and b are in the same row.
	for i := 0; i < len(mat); i++ {
		if AreAdjacent(mat[i], a, b) {
			return true
		}
	}
	// Check if a and b are in the same column.
	transposed := Transpose(mat)
	for i := 0; i < len(transposed); i++ {
		if AreAdjacent(transposed[i], a, b) {
			return true
		}
	}
	return false
}

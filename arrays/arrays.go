//Package for array and string utils
package arrays

//Checks if the input string is unique,
//don't have any repeated character
func IsUnique(phrase string) bool {
	var letters = map[byte]bool{0x00: false}

	for i := 0; i < len(phrase); i++ {
		if !letters[phrase[i]] {
			letters[phrase[i]] = true
		} else {
			return false
		}
	}
	return true
}

//Rotates an NxN matrix 90 degrees
func RotateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return nil
	}
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := 0; i < last; i++ {
			offset := i - first
			top := matrix[first][i]
			matrix[first][i] = matrix[last-offset][first]
			matrix[last-offset][first] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}
	return matrix
}

func nullifyColumn(matrix [][]int, column int) [][]int {
	for i := 0; i < len(matrix); i++ {
		matrix[i][column] = 0
	}
	return matrix
}

func nullifyRow(matrix [][]int, row int) [][]int {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[row][i] = 0
	}
	return matrix
}

func MatrixZero(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	var columns = make([]bool, m)
	var rows = make([]bool, n)
	for i := 0; i < n; i++ {
		for j := 0; i < m; j++ {
			if matrix[i][j] == 0 {
				columns[j] = true
				rows[i] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		if rows[i] {
			matrix = nullifyRow(matrix, i)
		}
	}

	for i := 0; i < m; i++ {
		if columns[i] {
			matrix = nullifyColumn(matrix, i)
		}
	}
	return matrix
}

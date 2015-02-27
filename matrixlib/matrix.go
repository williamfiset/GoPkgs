package matrixlib

import (
	"fmt"
	// "math"
)

type Matrix interface {
	Mul(m Matrix) Matrix
	Add(m Matrix) Matrix
}

type DenseMatrix struct {
	matrix [][]int
	rows   uint
	cols   uint
}

func NewMatrix(m [][]int) *DenseMatrix {
	var rows uint = uint(len(m))
	var columns uint = uint(len(m[0]))
	return &DenseMatrix{m, rows, columns}
}

// Create a Matrix of any dimension
func NewMatrixDim(rows uint, cols uint) *DenseMatrix {

	if rows > 0 && cols > 0 {

		newMatrix := make([][]int, rows)
		for i := uint(0); i < cols; i++ {

			newMatrix[i] = make([]int, cols)
			for j := uint(0); j < cols-1; j++ {
				newMatrix[i] = append(newMatrix[i], 0)
			}

		}

		return &DenseMatrix{newMatrix, rows, cols}
	}
	return nil
}

/*

	[1, 2] [1 3 0]  [a b e]
	[3, 1] [0 1 0]  [c d f]

	 2x2	2x3



*/

/*
func (a *DenseMatrix) Mul(b *DenseMatrix) *DenseMatrix {

	// Valid Multiplication
	if a.cols == b.rows {

		var newRowLen, newColLen uint = a.rows, b.cols
		fmt.Println(newRowLen, newColLen)
		ret := NewMatrixDim(newRowLen, newColLen)
		// ret := &DenseMatrix{nil, 1, 1}
		fmt.Println("WTF")
		return nil
		fmt.Println("WTF")

		fmt.Printf("R: %d, C: %d\n", newRowLen, newColLen)

		for r := uint(0); r < b.rows; r++ {
			for c := uint(0); c < b.cols; c++ {
				fmt.Println("HERE")
				k := 0
				for _, e := range a.matrix[r] {
					ret.matrix[r][c] += e * b.matrix[r][k] // Mul
					k++
				}
				// fmt.Println("")

			}
		}
		fmt.Println(ret.matrix)
		// for i := uint(0); i < newRowLen; i++ {
		// 	fmt.Println(row)
		// 	for j := 0; j < len(a.matrix[0]); j++ {
		// 		fmt.Println(row[i], a.matrix[i][j])
		// 		ret.matrix[j][i] += row[i] * a.matrix[i][j]
		// 	}
		// }

		return &ret
	}

	return nil
}

*/

func NewMatrixEmpty(r int, c int, fill int) [][]int {

	newMatrix := make([][]int, r)

	for row := 0; row < r; row++ {
		newMatrix[row] = make([]int, c)
		for col := 0; col < c; col++ {
			newMatrix[row][col] = fill
		}
	}

	return newMatrix

}

// [1, 2] [1 3 0]  [a b e]
// [3, 1] [0 1 0]  [c d f]

//  2x2	2x3

func Mul(a [][]int, b [][]int) [][]int {

	if a_col, b_row := len(a[0]), len(b); a_col == b_row {

		var rows, cols = len(a), len(b[0])
		newMatrix := NewMatrixEmpty(rows, cols, 0)

		for _, row := range a {
			for r, e := range row {
				for c := 0; c < cols; c++ {
					newMatrix[r][c] += e * b[r][c]
				}
			}
		}
		fmt.Println(newMatrix)

		return newMatrix
	}
	panic("Bad array dimension")

}

func (m *DenseMatrix) Add(a *DenseMatrix) (ret *DenseMatrix) {

	if m.cols != a.cols || m.rows != a.rows {
		return nil
	}

	ret = NewMatrixDim(m.cols, a.rows)

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.cols; j++ {
			ret.matrix[i][j] = m.matrix[i][j] + a.matrix[i][j]
		}
	}

	return ret
}

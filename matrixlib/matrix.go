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

// Multiplies two matrices together
func Mul(a [][]int, b [][]int) [][]int {

	var aRows, aCols = len(a), len(a[0])
	var bRows, bCols = len(b), len(b[0])

	if aCols == bRows {

		newMatrix := NewMatrixEmpty(aRows, bCols, 0)

		for r := 0; r < aRows; r++ {
			for c := 0; c < bCols; c++ {
				for k := 0; k < aCols; k++ {
					newMatrix[r][c] += a[r][k] * b[k][c]
				}
			}
		}
		return newMatrix
	}
	panic("Bad array dimension")

}

func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
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

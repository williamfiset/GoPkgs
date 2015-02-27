package matrixlib

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {

	a := NewMatrix([][]int{
		{1, 2, 3},
		{3, 4},
	})

	b := NewMatrix([][]int{
		{1, 2},
		{3, 4},
	})

	c := a.Add(b)

	if c != nil {
		t.Error("Supposed to be nil")
	}

	d := NewMatrix([][]int{
		{1, 2},
		{3, 4},
	})

	c = d.Add(b)

	if c == nil {
		t.Fatal("c is nil")
	}

	expected := [][]int{
		{2, 4},
		{6, 8},
	}

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			if expected[i][j] != c.matrix[i][j] {
				t.Error(expected[i][j], "!=", c.matrix[i][j])
			}
		}
	}

	// fmt.Println(c.matrix)

}

func Test_Mul(t *testing.T) {

	x := NewMatrix([][]int{
		{1, 2},
		{3, 1},
	})

	y := NewMatrix([][]int{
		{1, 3, 0},
		{0, 1, 0},
	})

	expected := [][]int{
		{1, 5, 0},
		{3, 10, 0},
	}
	expected2 := [][]int{
		{7, 4},
		{6, 7},
	}

	if !reflect.DeepEqual(expected, Mul(x.matrix, y.matrix)) {
		t.Error("Multiplication error")
	}
	// fmt.Println("BEFORE:", x.matrix)
	if !reflect.DeepEqual(expected2, Mul(x.matrix, x.matrix)) {
		t.Error("Multiplication error")
	}
	fmt.Println("FINISHED")

	// Panics as expected
	// z := NewMatrixEmpty(3, 1, 4)
	// Mul(x.matrix, z)

}

package main

import (
	"fmt"
	"github.com/william-fiset/mathutil"
	"github.com/william-fiset/matrix"
)

// import "github.com/william-fiset/mathutil/matrix"
// import "github.com/william-fiset/stringutil"

func check_mathutil() {

	fmt.Println(mathutil.Gcd(12, 18))
	fmt.Println(mathutil.Gcd(18, 12))

}

func check_matrixlib() {

	a := [][]int{
		{1, 2},
		{3, 4},
	}
	b := matrixlib.Mul(a, a)
	fmt.Println(b)

}

func main() {
	check_mathutil()
	check_matrixlib()
}

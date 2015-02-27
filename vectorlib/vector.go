package vectorlib

func NewVector(length int, fill int) []int {

	vector := make([]int, length)
	for i := 0; i < length; i++ {
		vector[i] = fill
	}
	return vector

}

// Remember, Dot Product Returns a Scaler!
func DotProduct(a []int, b []int) int {

	if len(a) == len(b) {
		tot := 0
		for i := 0; i < len(a); i++ {
			tot += a[i] * b[i]
		}
		return tot
	}

	panic("Cannot Dot Product: len(a) != len(b)")

}

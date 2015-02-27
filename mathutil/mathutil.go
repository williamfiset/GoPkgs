package mathutil

import "math"

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func IsPalindrome(str *string) bool {

	if str != nil {
		length := len(*str)
		for i := 0; i < length/2; i++ {
			if (*str)[i] != (*str)[length-i-1] {
				return false
			}
		}
	}
	return true
}

func IsPrime(n int64) bool {

	if n < 2 {
		return false
	} else if n == 2 || n == 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	root := int64(math.Sqrt(float64(n)) + 1.0)
	for f := int64(5); f < root; f += 6 {
		if n%f == 0 || n%(f+2) == 0 {
			return false
		}
	}
	return true
}

func DigitAtIndex(n, i int) int {
	return (n / int(math.Pow(10.0, float64(i)))) % 10
}

package vectorlib

import (
	"testing"
)

func Test_dotproduct(t *testing.T) {

	v0 := NewVector(2, 2)
	v1 := NewVector(2, 1)

	expected := 4

	if DotProduct(v0, v1) != expected {
		t.Error("Failed Dot Product, got:", DotProduct(v0, v1))
	}

}

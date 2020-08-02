package tuples

import (
	"testing"
)

func TestTuple(t *testing.T) {
	a := Tuple(4.3, -4.2, 3.1, 1.0)
	if a.x != 4.3 || a.y != -4.2 || a.z != 3.1 || a.w != 1.0 {
		t.Fail()
	}
}

func TestPoint(t *testing.T) {
	p := Point(4, -4, 3)
	if p != Tuple(4, -4, 3, 1) {
		t.Fail()
	}
}

func TestVector(t *testing.T) {
	v := Vector(4, -4, 3)
	if v != Tuple(4, -4, 3, 0) {
		t.Fail()
	}
}

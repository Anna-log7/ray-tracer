package tuples

import (
	"testing"
)

func TestTuple(t *testing.T) {
	a := Tuple(4.3, -4.2, 3.1, 1.0)
	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 1.0 {
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

func TestAdditon(t *testing.T) {
	tup1 := Tuple(3, -2, 5, 1)
	tup2 := Tuple(-2, 3, 1, 0)

	if AddTuples(tup1, tup2) != Tuple(1, 1, 6, 1) {
		t.Fail()
	}
}

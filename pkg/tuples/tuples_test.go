package tuples

import (
	"math"
	"testing"
)

func TestNewTuple(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 1.0)
	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 1.0 {
		t.Errorf("NewTuple creation error")
	}
}

func TestEqual(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 1.0)
	b := NewTuple(4.3, -4.2, 3.1, 1.0)
	c := NewTuple(-1, -4.2, 3.1, 1.0)

	if !Equal(a, b) || Equal(a, c) || Equal(b, c) {
		t.Errorf("NewTuple equal operator error")
	}
}

func TestNewPoint(t *testing.T) {
	point := NewPoint(4, -4, 3)
	if !Equal(point, NewTuple(4, -4, 3, 1)) {
		t.Errorf("NewPoint creation error")
	}
}

func TestNewVector(t *testing.T) {
	vector := NewVector(4, -4, 3)
	if !Equal(vector, NewTuple(4, -4, 3, 0)) {
		t.Errorf("NewVector creation error")
	}
}

func TestAdditon(t *testing.T) {
	tup1 := NewTuple(3, -2, 5, 1)
	tup2 := NewTuple(-2, 3, 1, 0)

	if !Equal(Add(tup1, tup2), NewTuple(1, 1, 6, 1)) {
		t.Errorf("Tupple addition error %v %v", tup1, tup2)
	}
}

func TestSubtractNewPoints(t *testing.T) {
	point1 := NewPoint(3, 2, 1)
	point2 := NewPoint(5, 6, 7)

	if !Equal(Subtract(point1, point2), NewVector(-2, -4, -6)) {
		t.Errorf("Subtract points error %v %v", point1, point2)
	}
}

func TestSubtractNewVectorFromNewPoint(t *testing.T) {
	point := NewPoint(3, 2, 1)
	vector := NewVector(5, 6, 7)

	if !Equal(Subtract(point, vector), NewPoint(-2, -4, -6)) {
		t.Errorf("Subtract vector from point error %v %v", point, vector)
	}
}

func TestSubtractNewVectors(t *testing.T) {
	vector1 := NewVector(3, 2, 1)
	vector2 := NewVector(5, 6, 7)

	if !Equal(Subtract(vector1, vector2), NewVector(-2, -4, -6)) {
		t.Errorf("Subtract vectors error %v %v", vector1, vector2)
	}
}

func TestNegateNewVectors(t *testing.T) {
	tup := NewTuple(1, -2, 3, -4)

	if !Equal(Negate(tup), NewTuple(-1, 2, -3, 4)) {
		t.Errorf("Negate tuple error %v", tup)
	}
}

func TestMultiply(t *testing.T) {
	tup := NewTuple(1, -2, 3, -4)
	scalar := 3.5

	if !Equal(Multiply(tup, scalar), NewTuple(3.5, -7, 10.5, -14)) {
		t.Errorf("NewTuple multiplication error %v %f", tup, scalar)
	}
}

func TestDivide(t *testing.T) {
	tup := NewTuple(1, -2, 3, -4)
	scalar := 2.0

	if !Equal(Divide(tup, scalar), NewTuple(.5, -1, 1.5, -2)) {
		t.Errorf("NewTuple division error %v %f", tup, scalar)
	}
}

func TestMagniture(t *testing.T) {
	vector1 := NewVector(1, 0, 0)
	const vectorError = "NewVector magnitude error %v"

	if Magnitude(vector1) != 1 {
		t.Errorf(vectorError, vector1)
	}

	vector2 := NewVector(0, 1, 0)

	if Magnitude(vector2) != 1 {
		t.Errorf(vectorError, vector2)
	}

	vector3 := NewVector(0, 0, 1)

	if Magnitude(vector3) != 1 {
		t.Errorf(vectorError, vector3)
	}

	vector4 := NewVector(1, 2, 3)

	if Magnitude(vector4) != math.Sqrt(14) {
		t.Errorf(vectorError, vector4)
	}

	vector5 := NewVector(-1, -2, -3)

	if Magnitude(vector5) != math.Sqrt(14) {
		t.Errorf(vectorError, vector5)
	}
}

func TestNormalization(t *testing.T) {
	const normalizationError = "Normalization error %v"
	vector1 := NewVector(4, 0, 0)

	if !Equal(Normalize(vector1), NewVector(1, 0, 0)) {
		t.Errorf(normalizationError, vector1)
	}

	vector2 := NewVector(1, 2, 3)

	if !Equal(Normalize(vector2), NewVector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))) {
		t.Errorf(normalizationError, vector2)
	}

	norm := Normalize(vector2)

	if Magnitude(norm) != 1 {
		t.Errorf("Magnitude of normalized vector error %v %v", vector2, norm)
	}
}

func TestDotProduct(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)

	if DotProduct(vector1, vector2) != 20 {
		t.Errorf("Dot product of vector error %v %v", vector1, vector2)
	}
}

func TestCrossProduct(t *testing.T) {
	const crossError = "Cross product of vector error %v %v"
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)

	if !Equal(CrossProduct(vector1, vector2), NewVector(-1, 2, -1)) {
		t.Errorf(crossError, vector1, vector2)
	}

	if !Equal(CrossProduct(vector2, vector1), NewVector(1, -2, 1)) {
		t.Errorf(crossError, vector1, vector2)
	}
}

package tuples

import (
	"math"
	"testing"
)

func TestTuple(t *testing.T) {
	a := Tuple(4.3, -4.2, 3.1, 1.0)
	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 1.0 {
		t.Errorf("Tuple creation error")
	}
}

func TestEqual(t *testing.T) {
	a := Tuple(4.3, -4.2, 3.1, 1.0)
	b := Tuple(4.3, -4.2, 3.1, 1.0)
	c := Tuple(-1, -4.2, 3.1, 1.0)

	if !Equal(a, b) || Equal(a, c) || Equal(b, c) {
		t.Errorf("Tuple equal operator error")
	}
}

func TestPoint(t *testing.T) {
	point := Point(4, -4, 3)
	if !Equal(point, Tuple(4, -4, 3, 1)) {
		t.Errorf("Point creation error")
	}
}

func TestVector(t *testing.T) {
	vector := Vector(4, -4, 3)
	if !Equal(vector, Tuple(4, -4, 3, 0)) {
		t.Errorf("Vector creation error")
	}
}

func TestAdditon(t *testing.T) {
	tup1 := Tuple(3, -2, 5, 1)
	tup2 := Tuple(-2, 3, 1, 0)

	if !Equal(Add(tup1, tup2), Tuple(1, 1, 6, 1)) {
		t.Errorf("Tupple addition error %v %v", tup1, tup2)
	}
}

func TestSubtractPoints(t *testing.T) {
	point1 := Point(3, 2, 1)
	point2 := Point(5, 6, 7)

	if !Equal(Subtract(point1, point2), Vector(-2, -4, -6)) {
		t.Errorf("Subtract points error %v %v", point1, point2)
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	point := Point(3, 2, 1)
	vector := Vector(5, 6, 7)

	if !Equal(Subtract(point, vector), Point(-2, -4, -6)) {
		t.Errorf("Subtract vector from point error %v %v", point, vector)
	}
}

func TestSubtractVectors(t *testing.T) {
	vector1 := Vector(3, 2, 1)
	vector2 := Vector(5, 6, 7)

	if !Equal(Subtract(vector1, vector2), Vector(-2, -4, -6)) {
		t.Errorf("Subtract vectors error %v %v", vector1, vector2)
	}
}

func TestNegateVectors(t *testing.T) {
	tup := Tuple(1, -2, 3, -4)

	if !Equal(Negate(tup), Tuple(-1, 2, -3, 4)) {
		t.Errorf("Negate tuple error %v", tup)
	}
}

func TestMultiply(t *testing.T) {
	tup := Tuple(1, -2, 3, -4)
	scalar := 3.5

	if !Equal(Multiply(tup, scalar), Tuple(3.5, -7, 10.5, -14)) {
		t.Errorf("Tuple multiplication error %v %f", tup, scalar)
	}
}

func TestDivide(t *testing.T) {
	tup := Tuple(1, -2, 3, -4)
	scalar := 2.0

	if !Equal(Divide(tup, scalar), Tuple(.5, -1, 1.5, -2)) {
		t.Errorf("Tuple division error %v %f", tup, scalar)
	}
}

func TestMagniture(t *testing.T) {
	vector1 := Vector(1, 0, 0)
	const vectorError = "Vector magnitude error %v"

	if Magnitude(vector1) != 1 {
		t.Errorf(vectorError, vector1)
	}

	vector2 := Vector(0, 1, 0)

	if Magnitude(vector2) != 1 {
		t.Errorf(vectorError, vector2)
	}

	vector3 := Vector(0, 0, 1)

	if Magnitude(vector3) != 1 {
		t.Errorf(vectorError, vector3)
	}

	vector4 := Vector(1, 2, 3)

	if Magnitude(vector4) != math.Sqrt(14) {
		t.Errorf(vectorError, vector4)
	}

	vector5 := Vector(-1, -2, -3)

	if Magnitude(vector5) != math.Sqrt(14) {
		t.Errorf(vectorError, vector5)
	}
}

func TestNormalization(t *testing.T) {
	const normalizationError = "Normalization error %v"
	vector1 := Vector(4, 0, 0)

	if !Equal(Normalize(vector1), Vector(1, 0, 0)) {
		t.Errorf(normalizationError, vector1)
	}

	vector2 := Vector(1, 2, 3)

	if !Equal(Normalize(vector2), Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))) {
		t.Errorf(normalizationError, vector2)
	}

	norm := Normalize(vector2)

	if Magnitude(norm) != 1 {
		t.Errorf("Magnitude of normalized vector error %v %v", vector2, norm)
	}
}

func TestDotProduct(t *testing.T) {
	vector1 := Vector(1, 2, 3)
	vector2 := Vector(2, 3, 4)

	if Dot(vector1, vector2) != 20 {
		t.Errorf("Dot product of vector error %v %v", vector1, vector2)
	}
}

func TestCrossProduct(t *testing.T) {
	const crossError = "Cross product of vector error %v %v"
	vector1 := Vector(1, 2, 3)
	vector2 := Vector(2, 3, 4)

	if !Equal(Cross(vector1, vector2), Vector(-1, 2, -1)) {
		t.Errorf(crossError, vector1, vector2)
	}

	if !Equal(Cross(vector2, vector1), Vector(1, -2, 1)) {
		t.Errorf(crossError, vector1, vector2)
	}
}

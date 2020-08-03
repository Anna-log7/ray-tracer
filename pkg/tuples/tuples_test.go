package tuples

import (
	"math"
	"testing"
)

func TestNewTuple(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 1.0)
	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 1.0 {
		t.Errorf("Tuple creation error")
	}
}

func TestEqual(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 1.0)
	b := NewTuple(4.3, -4.2, 3.1, 1.0)
	c := NewTuple(-1, -4.2, 3.1, 1.0)

	if !Equal(a, b) || Equal(a, c) || Equal(b, c) {
		t.Errorf("Tuple equal operator error")
	}
}

func TestNewPoint(t *testing.T) {
	point := NewPoint(4, -4, 3)
	expected := NewTuple(4, -4, 3, 1)

	if !Equal(point, expected) {
		t.Errorf("NewPoint creation error")
	}
}

func TestNewVector(t *testing.T) {
	vector := NewVector(4, -4, 3)
	expected := NewTuple(4, -4, 3, 0)

	if !Equal(vector, expected) {
		t.Errorf("Vector creation error")
	}
}

func TestAdditon(t *testing.T) {
	tup1 := NewTuple(3, -2, 5, 1)
	tup2 := NewTuple(-2, 3, 1, 0)

	expected := NewTuple(1, 1, 6, 1)
	result := Add(tup1, tup2)

	if !Equal(result, expected) {
		t.Errorf("Tuple addition error %v + %v Got: %v Expected: %v", tup1, tup2, result, expected)
	}
}

func TestSubtractPoints(t *testing.T) {
	point1 := NewPoint(3, 2, 1)
	point2 := NewPoint(5, 6, 7)

	expected := NewVector(-2, -4, -6)
	result := Subtract(point1, point2)

	if !Equal(result, expected) {
		t.Errorf("Subtract points error %v - %v Got: %v Expected: %v", point1, point2, result, expected)
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	point := NewPoint(3, 2, 1)
	vector := NewVector(5, 6, 7)

	expected := NewPoint(-2, -4, -6)
	result := Subtract(point, vector)

	if !Equal(result, expected) {
		t.Errorf("Subtract vector from point error %v - %v Got: %v Expected: %v", point, vector, result, expected)
	}
}

func TestSubtractVectors(t *testing.T) {
	vector1 := NewVector(3, 2, 1)
	vector2 := NewVector(5, 6, 7)

	expected := NewVector(-2, -4, -6)
	result := Subtract(vector1, vector2)

	if !Equal(result, expected) {
		t.Errorf("Subtract vectors error %v - %v Got: %v Expected: %v", vector1, vector2, result, expected)
	}
}

func TestNegateTuple(t *testing.T) {
	tup := NewTuple(1, -2, 3, -4)

	expected := NewTuple(-1, 2, -3, 4)
	result := tup.Negate()

	if !Equal(result, expected) {
		t.Errorf("Negate tuple error %v Got: %v Expected: %v", tup, result, expected)
	}
}

func TestMultiply(t *testing.T) {
	tup := NewTuple(1, -2, 3, -4)
	scalar := 3.5

	expected := NewTuple(3.5, -7, 10.5, -14)
	result := Multiply(tup, scalar)

	if !Equal(result, expected) {
		t.Errorf("Tuple multiplication error %v * %f Got: %v Expected: %v", tup, scalar, result, expected)
	}
}

func TestDivide(t *testing.T) {
	tup := NewTuple(1, -2, 3, -4)
	scalar := 2.0

	expected := NewTuple(.5, -1, 1.5, -2)
	result := Divide(tup, scalar)

	if !Equal(result, expected) {
		t.Errorf("Tuple division error %v / %f Got: %v Expected: %v", tup, scalar, result, expected)
	}
}

func TestMagniture(t *testing.T) {
	const vectorError = "Vector magnitude error %v Got: %f Expected: %f"
	vector1 := NewVector(1, 0, 0)

	expected1 := 1.0
	result1 := vector1.Magnitude()

	if !FloatEqual(result1, expected1) {
		t.Errorf(vectorError, vector1, result1, expected1)
	}

	vector2 := NewVector(0, 1, 0)
	result2 := vector2.Magnitude()

	if !FloatEqual(result2, expected1) {
		t.Errorf(vectorError, vector2, result2, expected1)
	}

	vector3 := NewVector(0, 0, 1)
	result3 := vector3.Magnitude()

	if !FloatEqual(result3, expected1) {
		t.Errorf(vectorError, vector3, result3, expected1)
	}

	vector4 := NewVector(1, 2, 3)
	expected4 := math.Sqrt(14)
	result4 := vector4.Magnitude()

	if !FloatEqual(result4, expected4) {
		t.Errorf(vectorError, vector4, result4, expected4)
	}

	vector5 := NewVector(-1, -2, -3)
	result5 := vector5.Magnitude()

	if !FloatEqual(result5, expected4) {
		t.Errorf(vectorError, vector5, result5, expected4)
	}
}

func TestNormalization(t *testing.T) {
	const normalizationError = "Normalization error %v Got: %v Expected: %v"
	vector1 := NewVector(4, 0, 0)
	expected1 := NewVector(1, 0, 0)
	result1 := vector1.Normalize()

	if !Equal(expected1, result1) {
		t.Errorf(normalizationError, vector1, result1, expected1)
	}

	vector2 := NewVector(1, 2, 3)
	expected2 := NewVector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))
	result2 := vector2.Normalize()

	if !Equal(expected2, result2) {
		t.Errorf(normalizationError, vector2, result2, expected2)
	}

	norm := vector2.Normalize()
	expected3 := 1.0
	result3 := norm.Magnitude()

	if !FloatEqual(expected3, result3) {
		t.Errorf("Magnitude of normalized vector error %v %v Got: %f Expected: %f", vector2, norm, result3, expected3)
	}
}

func TestDotProduct(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)
	expected := 20.0
	result := DotProduct(vector1, vector2)

	if !FloatEqual(expected, result) {
		t.Errorf("Dot product of vector error %v %v Got: %f Expected: %f", vector1, vector2, result, expected)
	}
}

func TestCrossProduct(t *testing.T) {
	const crossError = "Cross product of vector error %v %v Got: %v Expected: %v"
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)

	result1 := CrossProduct(vector1, vector2)
	expected1 := NewVector(-1, 2, -1)

	if !Equal(result1, expected1) {
		t.Errorf(crossError, vector1, vector2, result1, expected1)
	}

	result2 := CrossProduct(vector2, vector1)
	expected2 := NewVector(1, -2, 1)

	if !Equal(result2, expected2) {
		t.Errorf(crossError, vector1, vector2, result2, expected2)
	}
}

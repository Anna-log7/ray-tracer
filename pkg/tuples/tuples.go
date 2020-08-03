package tuples

import (
	"math"
)

// Tuple data type
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// NewTuple returns tuple
func NewTuple(x float64, y float64, z float64, w float64) Tuple {
	return Tuple{x, y, z, w}
}

// NewPoint returns tuple with w = 1
func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

// NewVector returns tuple with w = 0
func NewVector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

// FloatEqual checks if 2 floats are equal
func FloatEqual(val1 float64, val2 float64) bool {
	const EPSILLON = .00001
	return math.Abs(val1-val2) < EPSILLON
}

// Equal checks if 2 tuples are equal
func Equal(tup1 Tuple, tup2 Tuple) bool {
	if FloatEqual(tup1.X, tup2.X) && FloatEqual(tup1.Y, tup2.Y) && FloatEqual(tup1.Z, tup2.Z) && FloatEqual(tup1.W, tup2.W) {
		return true
	}
	return false
}

// Add adds 2 tuples together
func Add(tup1 Tuple, tup2 Tuple) Tuple {
	return NewTuple(
		tup1.X+tup2.X,
		tup1.Y+tup2.Y,
		tup1.Z+tup2.Z,
		tup1.W+tup2.W,
	)
}

// Subtract subracts 2 tuples
func Subtract(tup1 Tuple, tup2 Tuple) Tuple {
	return NewTuple(
		tup1.X-tup2.X,
		tup1.Y-tup2.Y,
		tup1.Z-tup2.Z,
		tup1.W-tup2.W,
	)
}

// Negate negates a tuple
func Negate(tup Tuple) Tuple {
	return NewTuple(
		-tup.X,
		-tup.Y,
		-tup.Z,
		-tup.W,
	)
}

// Multiply multiplies a tuple by a scalar
func Multiply(tup Tuple, val float64) Tuple {
	return NewTuple(
		tup.X*val,
		tup.Y*val,
		tup.Z*val,
		tup.W*val,
	)
}

// Divide divides a tuple by a scalar
func Divide(tup Tuple, val float64) Tuple {
	return NewTuple(
		tup.X/val,
		tup.Y/val,
		tup.Z/val,
		tup.W/val,
	)
}

// Magnitude calculates the magnitude of a tuple
func Magnitude(tup Tuple) float64 {
	squareSum := math.Pow(tup.X, 2) + math.Pow(tup.Y, 2) + math.Pow(tup.Z, 2) + math.Pow(tup.W, 2)
	return math.Sqrt(squareSum)
}

// Normalize normalizes a tuple
func Normalize(tup Tuple) Tuple {
	magnitude := Magnitude(tup)

	return NewTuple(
		tup.X/magnitude,
		tup.Y/magnitude,
		tup.Z/magnitude,
		tup.W/magnitude,
	)
}

// DotProduct computes the dot product of 2 tuples
func DotProduct(tup1 Tuple, tup2 Tuple) float64 {
	return (tup1.X*tup2.X +
		tup1.Y*tup2.Y +
		tup1.Z*tup2.Z +
		tup1.W*tup2.W)
}

// CrossProduct returns a cross product vector of 2 vectors
func CrossProduct(vector1 Tuple, vector2 Tuple) Tuple {
	return NewVector(
		vector1.Y*vector2.Z-vector1.Z*vector2.Y,
		vector1.Z*vector2.X-vector1.X*vector2.Z,
		vector1.X*vector2.Y-vector1.Y*vector2.X)
}

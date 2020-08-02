package tuples

import (
	"math"
)

// TupleStruct data type
type TupleStruct struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Tuple returns tuple
func Tuple(x float64, y float64, z float64, w float64) TupleStruct {
	return TupleStruct{x, y, z, w}
}

// Point returns tuple with w = 1
func Point(x float64, y float64, z float64) TupleStruct {
	return TupleStruct{x, y, z, 1}
}

// Vector returns tuple with w = 0
func Vector(x float64, y float64, z float64) TupleStruct {
	return TupleStruct{x, y, z, 0}
}

// FloatEqual checks if 2 floats are equal
func FloatEqual(val1 float64, val2 float64) bool {
	const EPSILLON = .00001
	return math.Abs(val1-val2) < EPSILLON
}

// Equal checks if 2 tuples are equal
func Equal(tup1 TupleStruct, tup2 TupleStruct) bool {
	if FloatEqual(tup1.X, tup2.X) && FloatEqual(tup1.Y, tup2.Y) && FloatEqual(tup1.Z, tup2.Z) && FloatEqual(tup1.W, tup2.W) {
		return true
	}
	return false
}

// Add adds 2 tuples together
func Add(tup1 TupleStruct, tup2 TupleStruct) TupleStruct {
	return Tuple(
		tup1.X+tup2.X,
		tup1.Y+tup2.Y,
		tup1.Z+tup2.Z,
		tup1.W+tup2.W,
	)
}

// Subtract subracts 2 tuples
func Subtract(tup1 TupleStruct, tup2 TupleStruct) TupleStruct {
	return Tuple(
		tup1.X-tup2.X,
		tup1.Y-tup2.Y,
		tup1.Z-tup2.Z,
		tup1.W-tup2.W,
	)
}

// Negate negates a tuple
func Negate(tup TupleStruct) TupleStruct {
	return Tuple(
		-tup.X,
		-tup.Y,
		-tup.Z,
		-tup.W,
	)
}

func Multiply(tup TupleStruct, val float64) TupleStruct {
	return Tuple(
		tup.X*val,
		tup.Y*val,
		tup.Z*val,
		tup.W*val,
	)
}

func Divide(tup TupleStruct, val float64) TupleStruct {
	return Tuple(
		tup.X/val,
		tup.Y/val,
		tup.Z/val,
		tup.W/val,
	)
}

func Magnitude(tup TupleStruct) float64 {
	squareSum := math.Pow(tup.X, 2) + math.Pow(tup.Y, 2) + math.Pow(tup.Z, 2) + math.Pow(tup.W, 2)
	return math.Sqrt(squareSum)
}

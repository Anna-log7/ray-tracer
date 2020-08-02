package tuples

import (
	"math"
	"reflect"
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

// FloatsEqual checks if 2 floats are equal
func FloatsEqual(val1 float64, val2 float64) bool {
	EPSILLON := .00001
	return math.Abs(val1-val2) < EPSILLON
}

// Equal checks if 2 tuples are equal
func Equal(tup1 TupleStruct, tup2 TupleStruct) bool {
	val1 := reflect.ValueOf(tup1)
	typeOfVal1 := val1.Type()

	for i := 0; i < val1.NumField(); i++ {
		field := typeOfVal1.Field(i).Name
		if !FloatsEqual(tup1[field], tup2[field]) {
			return false
		}
	}
	return true
}

func AddTuples(tup1 TupleStruct, tup2 TupleStruct) TupleStruct {
	x := tup1.x + tup2.x
	y := tup1.y + tup2.y
	z := tup1.z + tup2.z
	w := tup1.w + tup2.w

	return Tuple(x, y, z, w)
}

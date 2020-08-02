package tuples

type tuple struct {
	x float32
	y float32
	z float32
	w float32
}

// Tuple returns tuple
func Tuple(x float32, y float32, z float32, w float32) tuple {
	return tuple{x, y, z, w}
}

// Point returns tuple with w = 1
func Point(x float32, y float32, z float32) tuple {
	return Tuple(x, y, z, 1)
}

// Vector returns tuple with w = 0
func Vector(x float32, y float32, z float32) tuple {
	return Tuple(x, y, z, 0)
}

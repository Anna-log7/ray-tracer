package colors

import "ray-tracer/pkg/tuples"

// Color struct
type Color struct {
	Red, Green, Blue float64
}

// NewColor creates a new color struct
func NewColor(red float64, green float64, blue float64) Color {
	return Color{red, green, blue}
}

// Equal checks if 2 colors are equal
func Equal(color1 Color, color2 Color) bool {
	if tuples.FloatEqual(color1.Red, color2.Red) && tuples.FloatEqual(color1.Green, color2.Green) && tuples.FloatEqual(color1.Blue, color2.Blue) {
		return true
	}
	return false
}

// Add adds 2 colors and returns a color
func Add(color1 Color, color2 Color) Color {
	return NewColor(
		color1.Red+color2.Red,
		color1.Green+color2.Green,
		color1.Blue+color2.Blue,
	)
}

// Subtract subtracts color2 from color1 and returns a color
func Subtract(color1 Color, color2 Color) Color {
	return NewColor(
		color1.Red-color2.Red,
		color1.Green-color2.Green,
		color1.Blue-color2.Blue,
	)
}

// Multiply a color by a scalar
func Multiply(color Color, scalar float64) Color {
	return NewColor(
		color.Red*scalar,
		color.Green*scalar,
		color.Blue*scalar,
	)
}

// MultiplyColors returns the hadamard product of 2 colors
func MultiplyColors(color1 Color, color2 Color) Color {
	return NewColor(
		color1.Red*color2.Red,
		color1.Green*color2.Green,
		color1.Blue*color2.Blue,
	)
}

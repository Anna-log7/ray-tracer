package colors

import (
	"testing"
)

func TestColor(t *testing.T) {
	color := Color{-0.5, 0.4, 1.7}

	if color.Red != -0.5 || color.Green != 0.4 || color.Blue != 1.7 {
		t.Errorf("Color creation error %v", color)
	}
}

func TestEqual(t *testing.T) {
	color1 := Color{1, 2, 3}
	color2 := Color{1, 2, 3}
	color3 := Color{1, 9, 3}

	if !Equal(color1, color2) || Equal(color1, color3) || Equal(color2, color3) {
		t.Errorf("Color equality check error %v %v %v", color1, color2, color3)
	}
}

func TestAddColors(t *testing.T) {
	color1 := Color{0.9, 0.6, 0.75}
	color2 := Color{0.7, 0.1, 0.25}

	expected := Color{1.6, 0.7, 1.0}
	result := Add(color1, color2)

	if !Equal(result, expected) {
		t.Errorf("Color addition error %v + %v Got: %v Expected: %v", color1, color2, result, expected)
	}
}

func TestSubtractColors(t *testing.T) {
	color1 := Color{0.9, 0.6, 0.75}
	color2 := Color{0.7, 0.1, 0.25}

	expected := Color{0.2, 0.5, 0.5}
	result := Subtract(color1, color2)

	if !Equal(result, expected) {
		t.Errorf("Color subtraction error %v - %v Got: %v Expected: %v", color1, color2, result, expected)
	}
}

func TestMultiplyColorByScalar(t *testing.T) {
	color := Color{0.2, 0.3, 0.4}
	scalar := 2.0

	expected := Color{0.4, 0.6, 0.8}
	result := Multiply(color, scalar)

	if !Equal(result, expected) {
		t.Errorf("Color * scalar error %v * %v Got: %v Expected: %v", color, scalar, result, expected)
	}
}

func TestMultiplyColors(t *testing.T) {
	color1 := Color{1, 0.2, 0.4}
	color2 := Color{0.9, 1, 0.1}

	expected := Color{0.9, 0.2, 0.04}
	result := MultiplyColors(color1, color2)

	if !Equal(result, expected) {
		t.Errorf("Color * color error %v * %v Got: %v Expected: %v", color1, color2, result, expected)
	}
}

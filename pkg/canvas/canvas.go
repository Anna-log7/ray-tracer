package canvas

import (
	"ray-tracer/pkg/colors"
)

// Canvas type
type Canvas struct {
	Height, Width int
	Pixels        [][]colors.Color
}

// NewCanvas creates a new canvas struct with all pixels init'd to black
func NewCanvas(width int, height int) Canvas {
	pixels := make([][]colors.Color, width)
	for i := range pixels {
		pixels[i] = make([]colors.Color, height)
	}

	return Canvas{height, width, pixels}
}

// PixelAt returns the color of the pixel at x, y
func PixelAt(canvas Canvas, x int, y int) colors.Color {
	return canvas.Pixels[x][y]
}

// WritePixel writes a color to a pixel
func WritePixel(canvas *Canvas, x int, y int, color colors.Color) {
	if x < canvas.Width && y < canvas.Height {
		canvas.Pixels[x][y] = color
	}
}

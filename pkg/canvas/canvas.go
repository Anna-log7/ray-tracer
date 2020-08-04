package canvas

import (
	"fmt"
	"ray-tracer/internal/config"
	"ray-tracer/pkg/colors"
	"strings"
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
	canvas.Pixels[x][y] = color
}

// ToPpm returns the string representation of canvas pixels in ppm format
func (canvas Canvas) ToPpm() string {
	header := fmt.Sprintf("%s\n%d %d\n%d\n", config.PpmVersion, canvas.Width, canvas.Height, config.ColorRange)

	content := make([]string, 1)
	offset := 0
	for y := 0; y < canvas.Height; y++ {
		for x := 0; x < canvas.Width; x++ {
			if len(content[y+offset])+len(PixelAt(canvas, x, y).ToString())+2 > 70 {
				offset++
				content = append(content, "")
			}
			if x != 0 {
				content[y+offset] += " "
			}
			content[y+offset] += fmt.Sprintf("%s", PixelAt(canvas, x, y).ToString())
		}
		if y != canvas.Height-1 {
			content = append(content, "")
		}
	}

	joinedContent := strings.Join(content, "\n")
	// joinedContent += "\n"
	return header + joinedContent
}

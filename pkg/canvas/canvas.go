package canvas

import (
	"fmt"
	"math"
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

func adjustColor(colorValue float64, scalar int) int {
	scaledColor := int(math.Round(colorValue * float64(scalar)))
	switch {
	case scaledColor < 0:
		scaledColor = 0
	case scaledColor > scalar:
		scaledColor = scalar
	}
	return scaledColor
}

func isLineSpaceAvailable(line string, addToLine string) bool {
	if len(line)+2+len(addToLine) <= 70 {
		return true
	}
	return false
}

// ToPpm returns the string representation of canvas pixels in ppm format
func (canvas Canvas) ToPpm() string {
	header := fmt.Sprintf("%s\n%d %d\n%d\n", config.PpmVersion, canvas.Width, canvas.Height, config.ColorRange)

	content := make([]string, 1)
	offset := 0
	for y := 0; y < canvas.Height; y++ {
		for x := 0; x < canvas.Width; x++ {
			pixel := PixelAt(canvas, x, y)
			red := fmt.Sprintf("%d", adjustColor(pixel.Red, config.ColorRange))
			green := fmt.Sprintf("%d", adjustColor(pixel.Green, config.ColorRange))
			blue := fmt.Sprintf("%d", adjustColor(pixel.Blue, config.ColorRange))

			if isLineSpaceAvailable(content[y+offset], red) {
				if content[y+offset] != "" {
					content[y+offset] += " "
				}

				content[y+offset] += red
			} else {
				content = append(content, red)
				offset++
			}
			if isLineSpaceAvailable(content[y+offset], green) {
				if content[y+offset] != "" {
					content[y+offset] += " "
				}

				content[y+offset] += green
			} else {
				content = append(content, green)
				offset++
			}
			if isLineSpaceAvailable(content[y+offset], blue) {
				if content[y+offset] != "" {
					content[y+offset] += " "
				}

				content[y+offset] += blue
			} else {
				content = append(content, blue)
				offset++
			}
		}
		content = append(content, "")
	}

	joinedContent := strings.Join(content, "\n")
	joinedContent += "\n"
	return header + joinedContent
}

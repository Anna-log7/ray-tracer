package canvas

import (
	"fmt"
	"ray-tracer/pkg/colors"
	"strings"
	"testing"
)

func TestCanvas(t *testing.T) {
	canvas := NewCanvas(10, 20)

	if canvas.Width != 10 {
		t.Errorf("Canvas width init error %v", canvas.Width)
	}

	if canvas.Height != 20 {
		t.Errorf("Canvas height init error %v", canvas.Height)
	}

	expected := colors.NewColor(0, 0, 0)

	for y := 0; y < canvas.Height; y++ {
		for x := 0; x < canvas.Width; x++ {
			if !colors.Equal(canvas.Pixels[x][y], expected) {
				t.Errorf("Canvas not initialized with black screen error %v", canvas.Pixels[x][y])
			}
		}
	}
}

func TestWritingPixels(t *testing.T) {
	canvas := NewCanvas(10, 20)
	red := colors.NewColor(1, 0, 0)

	WritePixel(&canvas, 2, 3, red)
	result := PixelAt(canvas, 2, 3)

	if !colors.Equal(result, red) {
		t.Errorf("Writing color to pixel error Got: %v Expected: %v", result, red)
	}
}

func TestToPpmHeaders(t *testing.T) {
	canvas := NewCanvas(5, 3)
	ppm := canvas.ToPpm()

	expected := `P3
5 3
255`
	ppmByLines := strings.Split(ppm, "\n")

	headerLines := make([]string, 3)
	for i := 0; i < 3; i++ {
		headerLines[i] = ppmByLines[i]
	}
	result := strings.Join(headerLines, "\n")

	if expected != result {
		t.Errorf("PPM header creation error Got %s Expected %s", result, expected)
	}
}

func TestCanvasToPpmPixels(t *testing.T) {
	canvas := NewCanvas(5, 3)
	color1 := colors.NewColor(1.5, 0, 0)
	color2 := colors.NewColor(0, 0.5, 0)
	color3 := colors.NewColor(-0.5, 0, 1)

	WritePixel(&canvas, 0, 0, color1)
	WritePixel(&canvas, 2, 1, color2)
	WritePixel(&canvas, 4, 2, color3)
	ppm := canvas.ToPpm()

	expected := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255`

	ppmByLines := strings.Split(ppm, "\n")

	pixelLines := make([]string, 3)
	for i, j := 3, 0; i < 6; i, j = i+1, j+1 {
		pixelLines[j] = ppmByLines[i]
	}
	result := strings.Join(pixelLines, "\n")

	if expected != result {
		t.Errorf("PPM pixel creation error Got %s Expected %s", result, expected)
	}
}

func TestPpmLineLength(t *testing.T) {
	canvas := NewCanvas(10, 2)
	color := colors.NewColor(1, 0.8, 0.6)

	for y := 0; y < canvas.Height; y++ {
		for x := 0; x < canvas.Width; x++ {
			WritePixel(&canvas, x, y, color)
		}
	}

	ppm := canvas.ToPpm()

	expected := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153`

	ppmByLines := strings.Split(ppm, "\n")

	pixelLines := make([]string, 4)
	for i, j := 3, 0; i < 7; i, j = i+1, j+1 {
		pixelLines[j] = ppmByLines[i]
	}
	result := strings.Join(pixelLines, "\n")

	if expected != result {
		t.Errorf("PPM line length < 70 error Got %s Expected %s", result, expected)
	}
}

func TestPpmTerminateNewLine(t *testing.T) {
	canvas := NewCanvas(5, 3)

	ppm := canvas.ToPpm()
	expected := "\n"
	result := fmt.Sprintf("%b%b", ppm[len(ppm)-1], ppm[len(ppm)-2])

	if expected != result {
		t.Errorf("PPM not terminated with newline error")
	}
}

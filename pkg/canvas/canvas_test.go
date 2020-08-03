package canvas

import (
	"ray-tracer/pkg/colors"
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

	for x := 0; x < canvas.Width; x++ {
		for y := 0; y < canvas.Height; y++ {
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

package main

import (
	"fmt"
	"math"
	"os"
	"ray-tracer/pkg/canvas"
	"ray-tracer/pkg/colors"
	"ray-tracer/pkg/physics"
	"ray-tracer/pkg/tuples"
)

func toFile(fileContents string) {
	file, err := os.Create("./image.ppm")

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(file, fileContents)
	err = file.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("File contents written")
}

func main() {
	start := tuples.NewPoint(0, 1, 0)
	velocity := tuples.Multiply(tuples.NewVector(1, 1.8, 0).Normalize(), 11.25)
	proj := physics.Projectile{Position: start, Velocity: velocity}

	gravity := tuples.NewVector(0, -0.1, 0)
	wind := tuples.NewVector(-0.01, 0, 0)
	env := physics.Environment{Gravity: gravity, Wind: wind}

	color := colors.NewColor(0, 1.4, 1)
	c := canvas.NewCanvas(900, 550)

	for proj.Position.Y > 0 {
		canvas.WritePixel(&c, int(math.Round(proj.Position.X)), c.Height-int(math.Round(proj.Position.Y)), color)
		proj = physics.Tick(env, proj)
	}
	ppm := c.ToPpm()

	toFile(ppm)
}

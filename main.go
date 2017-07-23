package main

import (
	"flag"
	"fmt"
)

// Possible fractals
var fractals = []string{
	"mandelbrot",
	"sierpinski",
	"julia",
}

func main() {
	// Various flags we use pointers otherwise default value won't change
	size := flag.Int("size", 400, "Size of the fractal image in px")
	name := flag.String("name",
		"mandelbrot",
		fmt.Sprintf("Fractal Name, possible options are: %v", fractals))
	re := flag.Float64("re", 0.285, "Real part of the complex number for julia's set computation")
	im := flag.Float64("im", 0.0013, "Imaginary part of the complex number for julia's set computation")
	limit := flag.Int("limit", 200, "Limit of iteration to consider the sequence is bounded")

	// Parse flags
	flag.Parse()

	// Switch on fractal name that has been asked
	switch *name {
	case "mandelbrot":
		mandelbrot(float64(*size), float64(*limit))

	case "sierpinski":
		sierpinski(*size)

	case "julia":
		julia(float64(*size), float64(*limit), complex(*re, *im))

	default:
		fmt.Println("Sorry this fractal name isn't handled here")
	}
}

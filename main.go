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
	output := flag.String("output", "myfractal.jpg", "Name of the image file to output, format should be jpeg")
	colorized := flag.Bool("colorized", false, "If the output should be colorized (not possible for sierpinski ATM)")

	// Parse flags
	flag.Parse()

	// Switch on fractal name that have been asked
	switch *name {
	case "mandelbrot":
		mandelbrot(float64(*size), float64(*limit), *output, *colorized)

	case "sierpinski":
		sierpinski(*size, *output, *colorized)

	case "julia":
		julia(float64(*size), float64(*limit), complex(*re, *im), *output, *colorized)

	default:
		fmt.Println("Sorry this fractal name isn't handled here")
		return
	}
	fmt.Println("Computation Ended !")
}

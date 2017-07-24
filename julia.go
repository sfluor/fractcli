package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math/cmplx"
	"os"
	"sync"
)

// Check if complex number c is in Julia z0 set
func InJulia(z0, c complex128, n float64) (bool, float64) {
	z := z0
	for i := float64(0); i < n; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return false, i
		}
	}

	return true, n
}

// Create a julia img of the given size
func julia(size float64, limit float64, c complex128, output string, colorized bool) {
	// Create our image
	img := image.NewRGBA(image.Rect(0, 0, int(size), int(size)))
	// initialize image
	background := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)
	// We use goroutines to fasten the computation so we need a wait group
	var wg sync.WaitGroup
	// Wait for all columns to finish ( a go routine per column of pixels)
	wg.Add(int(size))

	mapColors := constructColorMap(limit, colorized)

	for x := float64(0); x < size; x++ {
		// Our go routine (we have to pass x as a value otherwise its value will change overtime)
		go func(img *image.RGBA, x float64) {
			// Wait group things
			defer wg.Done()
			// Check for our column
			for y := float64(0); y < size; y++ {
				_, gap := InJulia(complex(3*x/size-1.5, 3*y/size-1.5), c, limit)
				r, g, b := mapColors(gap)
				// Set the color of our pixel
				img.Set(int(x), int(y), color.RGBA{r, g, b, 255})
			}
		}(img, x)
	}
	// Wait for our goroutines to end
	wg.Wait()
	// Create the file where our image will be stored
	toimg, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer toimg.Close()
	// Register imagee
	jpeg.Encode(toimg, img, &jpeg.Options{jpeg.DefaultQuality})
}

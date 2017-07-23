package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var corner = [][2]float64{
	[2]float64{0, 0},
	[2]float64{0.5, math.Sqrt(3) / 2},
	[2]float64{1, 0},
}

// Return a random int between 0 and n-1
func random(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n)
}

// Return the middle of two points
func midpoint(p, q [2]float64) (float64, float64) {
	return 0.5 * (p[0] + q[0]), 0.5 * (p[1] + q[1])
}

// Create a sierpinski of the given size
func sierpinski(size int) {
	// Create our image
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	// initialize image
	background := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)

	// We spawn go routines
	var wg sync.WaitGroup
	wg.Add(size)
	for i := 0; i < size; i++ {
		x := 0.3
		y := 0.19
		go func(x, y float64) {
			defer wg.Done()
			for j := 0; j < size; j++ {
				img.Set(int(x*float64(size)), size-int(y*float64(size)), color.RGBA{255, 255, 255, 255})
				k := random(3)
				x, y = midpoint(corner[k], [2]float64{x, y})
			}
		}(x, y)
	}

	wg.Wait()
	// Create the file where our image will be stored
	toimg, _ := os.Create("sierpinski" + strconv.Itoa(size) + ".jpg")
	defer toimg.Close()
	// Register image
	jpeg.Encode(toimg, img, &jpeg.Options{jpeg.DefaultQuality})
}
package main

import colorful "github.com/lucasb-eyer/go-colorful"

type ColorMap func(float64) (uint8, uint8, uint8)

// Default HSL values for saturation and luminance
const (
	saturation = 0.5
	luminance  = 0.8
)

// return a function that Map the interval [0, limit] to [colors]
func constructColorMap(limit float64, colorized bool) ColorMap {
	if colorized {
		return func(x float64) (uint8, uint8, uint8) {
			r, g, b, _ := colorful.Hsl(x/limit-180, saturation, luminance*(1-x/limit)).RGBA()
			return uint8(r), uint8(g), uint8(b)
		}
	} else {
		return func(x float64) (uint8, uint8, uint8) {
			c := uint8(255 * x / limit)
			return c, c, c
		}
	}
}

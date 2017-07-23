package main

import "testing"

func benchmarkMendelbrot(size float64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		createImg(size, 200)
	}
}

// Basic benchmarks
func BenchmarkMendelbrot50(b *testing.B)  { benchmarkMendelbrot(50, b) }
func BenchmarkMendelbrot200(b *testing.B) { benchmarkMendelbrot(200, b) }
func BenchmarkMendelbrot400(b *testing.B) { benchmarkMendelbrot(400, b) }
func BenchmarkMendelbrot600(b *testing.B) { benchmarkMendelbrot(600, b) }
func BenchmarkMendelbrot800(b *testing.B) { benchmarkMendelbrot(800, b) }

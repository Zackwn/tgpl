package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	e6Lissajous()
}

func e6Lissajous() {
	// black and green
	var palette = []color.Color{
		color.RGBA{52, 56, 55, 1},
		color.RGBA{15, 155, 142, 1},
		color.RGBA{3, 113, 156, 1},
		color.RGBA{203, 1, 98, 1},
	}
	rand.Seed(time.Now().Unix())
	index := func() uint8 {
		x := rand.Intn(len(palette)-1) + 1
		return uint8(x)
	}
	file, _ := os.Create("gif.gif")
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index())
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(file, &anim) // NOTE: ignoring encoding errors
}

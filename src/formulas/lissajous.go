package formulas

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{135, 211, 124, 1},
	color.RGBA{255, 203, 5, 1},
	color.RGBA{210, 77, 87, 1},
	color.RGBA{107, 185, 240, 1},
	color.RGBA{118, 93, 105, 1},
	color.RGBA{77, 5, 232, 1},
	color.RGBA{244, 247, 118, 1},
}

const (
	blackIndex = 0
	greenIndex = 1
)

// Lissajous figures
func Lissajous(out io.Writer, c int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	cycles := c
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := 1
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
			if colorIndex == len(palette)-1 {
				colorIndex = 1
			} else {
				colorIndex++
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

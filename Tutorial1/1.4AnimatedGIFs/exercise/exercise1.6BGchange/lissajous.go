// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//exercise 1.5. green on black
var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.White,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 10    //delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	paletteLength := uint8(len(palette))
	for i := 0; i < nframes; i++ {
		var newPalette []color.Color
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		//this if is needed to slow background changing frequency
		if (uint8(i) % paletteLength) == 0 {
			//because the 0 index of the pallete is used for background collor for the image
			// we are going to shuffle the initial plalette so every image has different one
			shufftlePalette(palette)
		}
		//beacuse slice are passed by reference and they are not copied we must make
		// differen pallete for every image
		newPalette = make([]color.Color, len(palette))
		copy(newPalette, palette) // copy the initial pallete

		img := image.NewPaletted(rect, newPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := getRandomNumberDifferentFromX(paletteLength, 0)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//function wich helps the color index not to hit the background one
func getRandomNumberDifferentFromX(uperbound, x uint8) uint8 {
	for true {
		num := uint8((rand.Uint32() >> 24)) % uperbound
		if num != x {
			return num
		}
	}
	return 0
}

func shufftlePalette(palette []color.Color) {
	for i := len(palette) - 1; i > 0; i-- {
		j := rand.Uint32() % uint32(i)
		palette[i], palette[j] = palette[j], palette[i]
	}
}

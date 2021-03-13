// lissajous server exercise
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x26, 0x46, 0x53, 0xff},
	color.RGBA{0x2A, 0x9D, 0x8F, 0xff},
	color.RGBA{0xE9, 0xC4, 0x6A, 0xff},
	color.RGBA{0xF4, 0xA2, 0x61, 0xff},
	color.RGBA{0xE7, 0x6F, 0x51, 0xff},
}

// Draw a lissajous graphic to writer
func lissajous(out io.Writer, cycles int) {
	rand.Seed(time.Now().Unix())

	if cycles == 0 {
		cycles = 5
	}

	const (
		// cycles  = 5     // number of complete x oscillator revolutions
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
		colorIndex := uint8(rand.Intn(len(palette)-1)) + 1
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
		if err != nil {
			cycles = 5
		}

		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

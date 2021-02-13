package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

// TODO: is uint64 the best approach? Maybe stick with uint32 and batch the pixels?
// TODO: calculate median colour as well as mean. I find that mean is "muddy"
// TODO: comment your code you dingus

type colourCount struct {
	r     uint8
	g     uint8
	b     uint8
	count int
}

func getImageColour(filename string) color.Color {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // TODO: handle this error

	log.Printf("Successfully opened %s\n", filename)

	return modeColour(f)
}

func meanColour(f io.Reader) color.Color {
	i, fmtName, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully decoded %s\n", fmtName)

	var r, g, b uint64

	bounds := i.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pr, pg, pb, _ := i.At(x, y).RGBA()

			r += uint64(pr)
			g += uint64(pg)
			b += uint64(pb)
		}
	}

	d := uint64(bounds.Dy() * bounds.Dx())

	r /= d
	g /= d
	b /= d

	// TODO: should this be RGBA? Do I care about alpha?
	return color.NRGBA{R: uint8(r / 0x101), G: uint8(g / 0x101), B: uint8(b / 0x101), A: 255}
}

// TODO: Holy smokes what have I gotten myself into here. Color spaces, cielab, hilbert curve, what on earth?
func medianColour(f io.Reader) color.Color {
	return color.RGBA{}
}

func modeColour(f io.Reader) color.Color {
	i, fmtName, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully decoded %s\n", fmtName)

	bounds := i.Bounds()
	mColours := make(map[string]colourCount)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pr, pg, pb, _ := i.At(x, y).RGBA()

			r := uint8(pr / 0x101)
			g := uint8(pg / 0x101)
			b := uint8(pb / 0x101)

			colourString := fmt.Sprintf("%d %d %d", r, g, b)

			cc, ok := mColours[colourString]

			if ok {
				cc.count++
				mColours[colourString] = cc
			} else {
				cc.count = 1
				cc.r = r
				cc.g = g
				cc.b = b
				mColours[colourString] = cc
			}
		}
	}

	mode := getMaxCount(mColours)
	percentOfImage := float64(mode.count) / float64(bounds.Max.X * bounds.Max.Y)

	log.Printf("Max count: %d pixels, or %f%% of the image", mode.count, percentOfImage)

	return color.RGBA{R: mode.r, G: mode.g, B: mode.b}
}

func getMaxCount(m map[string]colourCount) colourCount {
	curMax := 0
	var curMaxColour colourCount

	for _, cc := range m {
		if cc.count > curMax {
			curMax = cc.count
			curMaxColour = cc
		}
	}

	return curMaxColour
}

func main() {
	filename := "cilantro.jpg"
	colour := getImageColour(filename)

	fmt.Printf("Mode colour of %s: %v\n", filename, colour)
}

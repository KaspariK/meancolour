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

// TODO: handle more than jpeg
// TODO: is uint64 the best approach? Maybe stick with uint32 and batch the pixels?
// TODO: calculate median colour as well as mean. I find that mean is "muddy"
// TODO: calculate mode colour. May as well do all of them

func getImageColour(filename string) color.Color{
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Printf("Successfully opened %s\n", filename)

	return meanColour(f)
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

	return color.NRGBA{R: uint8(r / 0x101), G: uint8(g / 0x101), B: uint8(b / 0x101), A: 255}
}

func main() {
	filename := "cowboy.png"
	colour := getImageColour(filename)

	fmt.Printf("Mean colour of %s: %v\n", filename, colour)
}
package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	_ "image/jpeg"
)

// TODO: handle more than jpeg
// TODO: unit testing. Possibly restructure to avoid file i/o
// TODO: is uint64 the best approach? Maybe stick with uint32 and batch the pixels?
// TODO: calculate median colour as well as mean. I find that mean is "muddy"
// TODO: calculate mode colour. May as well do all of them

func Image() {
	f, err := os.Open("white.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s successfully decoded\n", fmtName)

	colour := MeanColour(img)

	fmt.Printf("Mean image colour: %v\n", colour)
}

func MeanColour(i image.Image) color.Color {
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
	Image()
}
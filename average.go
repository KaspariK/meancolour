package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	_ "image/jpeg"
)

func Image() {
	f, err := os.Open("red.jpg")
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
	var r, g, b uint32

	bounds := i.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pr, pg, pb, _ := i.At(x, y).RGBA()

			r += pr
			g += pg
			b += pb
		}
	}

	d := uint32(bounds.Dy() * bounds.Dx())

	r /= d
	g /= d
	b /= d

	return color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func main() {
	Image()
}
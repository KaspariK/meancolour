package main

import (
	"image"
	"image/color"
	"os"
	"testing"
)

func TestGetImage(t *testing.T) {
	tests := map[string]struct{
		input string
		want image.Image
	}{
		"JPEG": {input:},
	}
}

func TestMeanColour(t *testing.T) {
	tests := map[string]struct{
		input image.Image
		want color.Color
	}{}
}
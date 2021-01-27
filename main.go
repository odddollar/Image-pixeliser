package main

import (
	"image"
	"image/png"
)

func main() {
	// create image decoder
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	// get number of desired colour blocks
	desiredX := 8
	desiredY := 8

	// open image and get width and height
	workingImage, workingImageWidth, workingImageHeight := readImage("testImage.png")

	// resize to temporary image of size that fits number of desired blocks closest to current width

	// get pixels from temporary image

	// group pixels together and average

	// write each pixel to output file

}

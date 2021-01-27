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
	newWidth, newHeight := getClosestDimensions(desiredX, desiredY, workingImageWidth, workingImageHeight)

	// get pixels from image, cropping to set dimensions
	pixels := getImagePixels(workingImage, newWidth, newHeight)

	// group pixels together and average
	pixels = groupPixels(pixels, desiredX, desiredY)

	// write each pixel to output file

}

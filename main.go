package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	// create image decoder
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("fpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)

	// create argparser
	parser := argparse.NewParser("Image pixeliser", "Converts a given image to a specified number of colour blocks on the x and y axes")
	file := parser.String("f", "file", &argparse.Options{Required: true, Help: "The path of the image file to convert"})
	width := parser.Int("x", "width", &argparse.Options{Required: true, Help: "The number of colour blocks to have along the x axis"})
	height := parser.Int("y", "height", &argparse.Options{Required: true, Help: "The number of colour blocks to have along the y axis"})

	// run argparser
	if err := parser.Parse(os.Args); err != nil {
		fmt.Println(parser.Usage(err))
	}

	// get number of desired colour blocks
	desiredX := *width
	desiredY := *height

	// open image and get width and height
	workingImage, workingImageWidth, workingImageHeight := readImage(*file)

	// resize to temporary image of size that fits number of desired blocks closest to current width
	newWidth, newHeight := getClosestDimensions(desiredX, desiredY, workingImageWidth, workingImageHeight)

	// get pixels from image, cropping to set dimensions
	pixels := getImagePixels(workingImage, newWidth, newHeight)

	// group pixels together and average
	pixels = groupPixels(pixels, desiredX, desiredY)

	// write each pixel to output file
	createImage(pixels)
}

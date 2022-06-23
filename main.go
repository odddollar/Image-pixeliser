package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	// create image decoder
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("fpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)

	// create argparser
	parser := argparse.NewParser("Image pixeliser", "Converts a given image to a specified number of colour blocks on the x and y axes")
	file := *parser.String("f", "file", &argparse.Options{Required: true, Help: "The path of the image file to convert"})
	desiredX := *parser.Int("x", "width", &argparse.Options{Required: true, Help: "The number of colour blocks to have along the x axis"})
	desiredY := *parser.Int("y", "height", &argparse.Options{Required: true, Help: "The number of colour blocks to have along the y axis"})
	output := *parser.String("o", "output", &argparse.Options{Required: false, Help: "The output directory/filename of the processed image"})

	// run argparser
	if err := parser.Parse(os.Args); err != nil {
		log.Fatal(parser.Usage(err))
	}

	// open image and get width and height
	workingImage, workingImageWidth, workingImageHeight := readImage(file)

	// resize to temporary image of size that fits number of desired blocks closest to current width
	newWidth, newHeight := getClosestDimensions(desiredX, desiredY, workingImageWidth, workingImageHeight)

	// get pixels from image, cropping to set dimensions
	pixels := getImagePixels(workingImage, newWidth, newHeight)

	// group pixels together and average
	pixels = groupPixels(pixels, desiredX, desiredY)

	// write each pixel to output file
	createImage(pixels, output)
}

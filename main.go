package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

type Pixel struct {
	r, g, b int
}

func main() {
	// get number of desired colour blocks

	// open image

	// resize to temporary image of size that fits number of desired blocks closest to current width

	// get pixels from temporary image

	// group pixels together and average

	// write each pixel to output file

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	fmt.Println(readImagePixels(readImage("testImage.png")))
}

func readImagePixels(image image.Image) [][]Pixel {
	// create pixel array
	var pixels [][]Pixel

	// get image bounds/size
	width, height := image.Bounds().Max.X, image.Bounds().Max.Y

	// iterate through height of image
	for y := 0; y < height; y++ {
		// create temporary row variable to store values of individual row
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(image.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	// return pixel array
	return pixels
}

func readImage(path string) image.Image {
	// open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()

	// decode image to array
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// return image
	return img
}

func rgbaToPixel(r uint32, g uint32, b uint32, _ uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257)}
}

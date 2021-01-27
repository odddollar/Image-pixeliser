package main

import (
	"image"
	"log"
	"os"
)

type Pixel struct {
	R, G, B int
}

func getImagePixels(image image.Image, width, height int) [][]Pixel {
	var pixels [][]Pixel

	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, convertToRGBA(image.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels
}

func convertToRGBA(r, g, b, _ uint32) Pixel {
	return Pixel{
		R: int(r / 257),
		G: int(g / 257),
		B: int(b / 257),
	}
}

func readImage(path string) (image.Image, int, int) {
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

	// return image, width and height
	return img, img.Bounds().Max.X, img.Bounds().Max.Y
}

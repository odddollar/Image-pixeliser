package main

import (
	"image"
	"log"
	"os"
)

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

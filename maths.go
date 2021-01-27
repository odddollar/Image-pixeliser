package main

func groupPixels(pixels [][]Pixel, desiredX, desiredY int) [][]Pixel {
	// get pixel width and height of colour blocks
	blockWidth := len(pixels[0]) / desiredX
	blockHeight := len(pixels) / desiredY

	// iterate through colour blocks
	for y := 0; y < len(pixels); y += blockHeight {
		for x := 0; x < len(pixels[0]); x += blockWidth {
			// create counters to find averages
			var totalR, totalG, totalB int
			var pixelCount int
			var averageR, averageG, averageB int

			// iterate through pixels in block, adding to total pixel count to find average
			for blockY := y; blockY < y+blockHeight; blockY++ {
				for blockX := x; blockX < x+blockWidth; blockX++ {
					totalR += pixels[blockY][blockX].R
					totalG += pixels[blockY][blockX].G
					totalB += pixels[blockY][blockX].B
					pixelCount++
				}
			}

			// find average rgb value of colour block
			averageR = totalR / pixelCount
			averageG = totalG / pixelCount
			averageB = totalB / pixelCount

			// set each pixel in colour block to average
			for blockY := y; blockY < y+blockHeight; blockY++ {
				for blockX := x; blockX < x+blockWidth; blockX++ {
					pixels[blockY][blockX].R = averageR
					pixels[blockY][blockX].G = averageG
					pixels[blockY][blockX].B = averageB
				}
			}
		}
	}

	// return average pixel array
	return pixels
}

func getClosestDimensions(desiredX, desiredY, width, height int) (int, int) {
	var resizedWidth, resizedHeight int

	// while i is less than the width
	for i := 0; i <= width; {
		if i+desiredX > width { // if not possible to step up another level, return current level
			resizedWidth = i
			break
		} else { // check if can step up another level, step up by that level
			i += desiredX
		}
	}

	// do same as above but with height
	for i := 0; i <= height; {
		if i+desiredY > height {
			resizedHeight = i
			break
		} else {
			i += desiredY
		}
	}

	// return dimensions that are divisible by the desired number of colour blocks, closest to the the current image dimensions
	return resizedWidth, resizedHeight
}

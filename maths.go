package main

func getClosestDimensions(desiredX, desiredY, width, height int) (int, int) {
	var resizedWidth, resizedHeight int

	// while i is less than the width
	for i := 0; i < width; {
		if i+desiredX < width { // check if can step up another level, step up by that level
			i += desiredX
		} else { // if not possible to step up another level, return current level
			resizedWidth = i
			break
		}
	}

	// do same as above but with height
	for i := 0; i < height; {
		if i+desiredY < height {
			i += desiredY
		} else {
			resizedHeight = i
			break
		}
	}

	// return dimensions that are divisible by the desired number of colour blocks, closest to the the current image dimensions
	return resizedWidth, resizedHeight
}

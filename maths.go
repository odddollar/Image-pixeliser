package main

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

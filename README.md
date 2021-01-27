# Image pixeliser

Command-line program to convert a given image into a specified number of colour blocks. 

The image is cropped to allow for the number of blocks to be divided evenly into it's dimensions (while still keeping as much of the image's original size as possible), then the pixel values are grouped together into blocks and averaged. Each pixel in this colour block is set to this average value. The output image is then written to disk.

Input: Image-pixeliser -f "Go logo.png" -x 16 -y 16

![Go logo](examples\Go logo.png) 

Output:

![Go logo pixelised](examples/Go logo pixelised.jpg)
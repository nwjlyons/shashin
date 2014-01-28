package main

import (
	"flag"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	width     = flag.Uint("w", 0, "Width to resize image to.")
	height    = flag.Uint("h", 0, "Height to resize image to.")
	grayscale = flag.Bool("g", false, "Convert image to grayscale.")
)

var usage = `Usage: shashin [options...] /file/path/to/image.jpg

Options:
  -w Width to resize image to.
  -h Height to resize image to.
  -g Covert image to grayscale.
`

func main() {
	flag.Parse()

	// Check command line argument is present
	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Image file path missing from command line arguments.\n\n")
		fmt.Fprintf(os.Stderr, usage+"\n")
		os.Exit(1)
	}

	imageFilepath := flag.Args()[0]

	// Open image
	imageFile, err := os.Open(imageFilepath)
	if err != nil {
		log.Fatal(err)
	}

	// Decode image to image.Image object.
	img, format, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	imageFile.Close()

	outputImage := resize.Resize(*width, *height, img, resize.Lanczos3)

	bounds := outputImage.Bounds()

	// Get pixel dimensions of image
	w, h := bounds.Max.X, bounds.Max.Y

	if *grayscale {
		// Create new grayscale image with the same dimensions as the old image.
		grayImage := image.NewGray(bounds)
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				// Get pixel colour at position x, y.
				oldColour := outputImage.At(x, y)

				// Set pixel colour at position x,y on new image.
				grayImage.Set(x, y, oldColour)
			}
		}
		outputImage = grayImage
	}

	extension := filepath.Ext(imageFilepath)
	imageFleName := strings.TrimRight(imageFilepath, extension)

	var gray string

	// If grayscale image add to output filename.
	if *grayscale {
		gray = "-grayscale"
	}

	outputFileName := fmt.Sprintf("%s-%dx%d%s%s", imageFleName, w, h, gray, extension)

	// Create file to save image to.
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	switch format {
	case "png":
		png.Encode(outputFile, outputImage)
	default:
		// jpeg by default
		jpeg.Encode(outputFile, outputImage, nil)
	}

}

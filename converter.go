package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/chai2010/webp"
)

func main() {
	inputPath := "./image.jpg"

	// Open JPG file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Decode JPG image
	img, err := jpeg.Decode(inputFile)
	if err != nil {
		fmt.Println("Error decoding JPEG:", err)
		return
	}

	rgbaImg := toRGBA(img)

	writeWebpImage("public/images/spread.webp", rgbaImg)
}

// toRGBA converts any image.Image to *image.RGBA
func toRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
	return dst
}

func writeWebpImage(outputPath string, img image.Image) {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	err = webp.Encode(outputFile, img, &webp.Options{Lossless: false, Quality: 75})
	if err != nil {
		fmt.Println("Error encoding to WebP:", err)
		return
	}
}

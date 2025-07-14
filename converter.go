package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"strconv"

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

	const numberOfTilesPerSide = 1

	dx := img.Bounds().Dx() / numberOfTilesPerSide
	dy := img.Bounds().Dy() / numberOfTilesPerSide

	for i := range numberOfTilesPerSide {
		for j := range numberOfTilesPerSide {
			outputPath := "output/" + strconv.Itoa(i) + "_" + strconv.Itoa(j) + ".webp"
			cropped := cropImage(rgbaImg, image.Rect(dx*i, dy*j, dx*(i+1), dy*(j+1)))
			writeWebpImage(outputPath, cropped)
		}
	}
}

// toRGBA converts any image.Image to *image.RGBA
func toRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
	return dst
}

// cropImage crops an image to the given rectangle.
func cropImage(src *image.RGBA, crop image.Rectangle) image.Image {
	crop = crop.Intersect(src.Bounds())
	return src.SubImage(crop)
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

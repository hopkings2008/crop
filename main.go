package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/artyom/smartcrop"
)

func main() {
	file := os.Args[1:]
	fi, err := os.Open(file[0])
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	img, _, err := image.Decode(fi)
	if err != nil {
		log.Fatal(err)
	}

	topCrop, err := smartcrop.Crop(img, 250, 250)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("best crop is", topCrop)

	type subImager interface {
		SubImage(image.Rectangle) image.Image
	}
	if si, ok := img.(subImager); ok {
		cr := si.SubImage(topCrop)
		fmt.Printf("cropped image dimensions are %d x %d\n", cr.Bounds().Dx(), cr.Bounds().Dy())
		writeImageToPng(cr, "./result.png")
	}
	// Output:
	// best crop is (59,0)-(486,427)
	// cropped image dimensions are 427 x 427
}

func writeImageToJpeg(img image.Image, name string) {
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		panic(err)
	}
	fso, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer fso.Close()

	jpeg.Encode(fso, img, &jpeg.Options{Quality: 100})
}

func writeImageToPng(img image.Image, name string) {
	fso, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer fso.Close()

	png.Encode(fso, img)
}

package utils

import (
	"image"
	"image/png"
	"log"
	"os"
)


func SaveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(f, img)
	f.Close()
}

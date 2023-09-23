package utils

import (
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
)


func SaveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}


func RandInterval(min, max float64) float64 {
	return min + (max - min) * rand.Float64()
}

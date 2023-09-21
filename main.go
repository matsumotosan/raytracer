package main

import (
	"fmt"
	"image"

	"raytracer/camera"
	"raytracer/geometry"
	"raytracer/utils"

	"github.com/schollz/progressbar/v3"
)


func main() {
	fmt.Println("Raytracing in Go")

	// Initialize camera
	cam := camera.Camera{
		Center:      geometry.Vec3{0, 0, 0},
		AspectRatio: 16.0 / 9.0,
		ImageWidth:  400,
	}

	cam.Initialize()

	image_width := cam.ImageWidth
	image_height := cam.ImageHeight
	pixel00_loc := cam.Pixel_00_loc
	pixel_delta_u := cam.Pixel_delta_u
	pixel_delta_v := cam.Pixel_delta_v

	// Create world
	world := geometry.World{}
	world.Add(geometry.Sphere{Center: geometry.Vec3{0,      0, -1}, Radius: 0.5})
	world.Add(geometry.Sphere{Center: geometry.Vec3{0, -100.5, -1}, Radius: 100})

	img := image.NewRGBA(image.Rect(0, 0, image_width, image_height))
	bar := progressbar.NewOptions(image_width * image_height)

	for i := 0; i < image_width; i++ {
		for j := 0; j < image_height; j++ {

			// Calculate next pixel center
			pixel_center := pixel00_loc.Add(pixel_delta_u.MulS(float64(i)))
			pixel_center = pixel_center.Add(pixel_delta_v.MulS(float64(j)))

			// Calculate ray to pixel center
			ray_vec := pixel_center.Sub(cam.Center)

			// Calculate ray color
			ray := geometry.Ray{Orig: cam.Center, Dir: ray_vec}
			color := cam.RayColor(ray, world)

			// Write color to pixel
			img.Set(i, j, color)

			bar.Add(1)
		}
	}

	utils.SaveImage("./test.png", img)
}

package main

import (
	"fmt"
	"image"
	"image/color"

	"raytracer/geometry"
	"raytracer/utils"

	"github.com/schollz/progressbar/v3"
)


func main() {
	fmt.Println("Raytracing in Go")

	// Calculate image height
	aspect_ratio := 16.0 / 9.0
	image_width := 400
	image_height := int(float64(image_width) / aspect_ratio)

	// Create world
	world := geometry.HittableList{}
	world.Add(geometry.Sphere{Center: geometry.Vec3{0,      0, -1}, Radius: 0.5})
	world.Add(geometry.Sphere{Center: geometry.Vec3{0, -100.5, -1}, Radius: 100})

	// Camera settings
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * float64(image_width) / float64(image_height)
	camera_center := geometry.Vec3{0, 0, 0}

	// Calculate total width and height of viewport
	viewport_u := geometry.Vec3{viewport_width, 0, 0}
	viewport_v := geometry.Vec3{0, -viewport_height, 0}

	// Calculate displacement size for horizontal and vertical vectors
	pixel_delta_u := viewport_u.DivS(float64(image_width))
	pixel_delta_v := viewport_v.DivS(float64(image_height))

	// Caluclate position of top left corner
	viewport_upper_left := camera_center.Sub(geometry.Vec3{0, 0, focal_length})
	viewport_upper_left = viewport_upper_left.Sub(viewport_u.DivS(2))
	viewport_upper_left = viewport_upper_left.Sub(viewport_v.DivS(2))

	// Calculate center of upper left pixel
	pixel00_loc := viewport_upper_left
	pixel00_loc = pixel00_loc.Add(pixel_delta_u.MulS(0.5))
	pixel00_loc = pixel00_loc.Add(pixel_delta_v.MulS(0.5))

	img := image.NewRGBA(image.Rect(0, 0, image_width, image_height))
	bar := progressbar.NewOptions(image_width * image_height)

	for i := 0; i < image_width; i++ {
		for j := 0; j < image_height; j++ {

			// Calculate next pixel center
			pixel_center := pixel00_loc.Add(pixel_delta_u.MulS(float64(i)))
			pixel_center = pixel_center.Add(pixel_delta_v.MulS(float64(j)))

			// Calculate ray to pixel center
			ray_vec := pixel_center.Sub(camera_center)

			// Calculate ray color
			ray := geometry.Ray{Orig: camera_center, Dir: ray_vec}
			color := RayColor(ray, world)

			// Write color to pixel
			img.Set(i, j, color)

			bar.Add(1)
		}
	}

	utils.SaveImage("./test.png", img)
}


func RayColor(ray geometry.Ray, world geometry.HittableList) color.RGBA {
	record := geometry.HitRecord{}

	if world.Hit(ray, geometry.Interval{Min: 0, Max: 999999}, &record) {
		return color.RGBA{
			uint8(255 / 2 * (record.Normal[0] + 1)),
			uint8(255 / 2 * (record.Normal[1] + 1)),
			uint8(255 / 2 * (record.Normal[2] + 1)),
			255,
		}
	}

	unit_ray := ray.Dir.GetUnit()
	a := 0.5 * (unit_ray[1] + 1.0)
	c := geometry.Vec3{1.0, 1.0, 1.0}.MulS(1.0 - a)
	c = c.Add(geometry.Vec3{0.5, 0.7, 1.0}.MulS(a))

	return color.RGBA{
		uint8(255 * c[0]),
		uint8(255 * c[1]),
		uint8(255 * c[2]),
		255,
	}
}

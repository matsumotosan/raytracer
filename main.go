package main

import (
	"fmt"

	"raytracer/camera"
	"raytracer/geometry"
	"raytracer/utils"
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

	// Create world
	world := geometry.World{}
	world.Add(geometry.Sphere{Center: geometry.Vec3{0,      0, -1}, Radius: 0.5})
	world.Add(geometry.Sphere{Center: geometry.Vec3{3,      2, -2}, Radius: 1})
	world.Add(geometry.Sphere{Center: geometry.Vec3{-3,      2, -2}, Radius: 1})
	world.Add(geometry.Sphere{Center: geometry.Vec3{0, -100.5, -1}, Radius: 100})

	// Take a picture
	img := cam.Render(&world)
	utils.SaveImage("./test.png", img)
}

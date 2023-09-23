package main

import (
	"fmt"

	"raytracer/camera"
	"raytracer/geometry"
	"raytracer/utils"
)


func main() {
	fmt.Println("Raytracing in Go")

	cam := camera.Camera{
		Center:          geometry.Vec3{0, 0, 0},
		AspectRatio:     16.0 / 9.0,
		ImageWidth:      400,
		SamplesPerPixel: 100,
	}

	cam.Initialize()

	world := geometry.World{}
	world.Add(geometry.Sphere{Center: geometry.Vec3{0,      0, -1}, Radius: 0.5})
	world.Add(geometry.Sphere{Center: geometry.Vec3{0, -100.5, -1}, Radius: 100})

	img := cam.Render(&world)
	utils.SaveImage("./test.png", img)
}

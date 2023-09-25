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
		ImageHeight:     0,
		SamplesPerPixel: 100,
		MaxBounces:      10,
	}

	cam.Initialize()

	world := geometry.World{}

	// Hollow glass sphere
	world.Add(geometry.Sphere{
		Center: geometry.Vec3{-1, 0, -1},
		Radius: -0.4,
		Material: &geometry.Dielectric { RefractiveIndex: 1.5 },
	})

	world.Add(geometry.Sphere{
		Center: geometry.Vec3{-1, 0, -1},
		Radius: 0.5,
		Material: &geometry.Dielectric { RefractiveIndex: 1.5 },
	})

	// Matte
	world.Add(geometry.Sphere{
		Center: geometry.Vec3{ 0, 0, -1},
		Radius: 0.5,
		Material: &geometry.Lambertian { Albedo: geometry.Vec3{0.1, 0.2, 0.5}},
	})

	// Metallic
	world.Add(geometry.Sphere{
		Center: geometry.Vec3{ 1, 0, -1},
		Radius: 0.5,
		Material: &geometry.Metal { Albedo: geometry.Vec3{0.8, 0.6, 0.2}, Fuzz: 0.0},
	})

	// Matte world
	world.Add(geometry.Sphere{
		Center: geometry.Vec3{0, -100.5, -1},
		Radius: 100,
		Material: &geometry.Lambertian{ Albedo: geometry.Vec3{ 0.8, 0.3, 0.2}},
	})

	img := cam.Render(&world)
	utils.SaveImage("./example.png", img)
}

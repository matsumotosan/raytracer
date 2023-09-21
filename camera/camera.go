package camera

import (
	"image"
	"image/color"

	"raytracer/geometry"
	"github.com/schollz/progressbar/v3"
)


type Camera struct {
	Center geometry.Vec3
	AspectRatio float64
	ImageWidth int
	ImageHeight int
	Pixel_00_loc geometry.Vec3
	Pixel_delta_u geometry.Vec3
	Pixel_delta_v geometry.Vec3
}


func (cam *Camera) Initialize() {
	cam.ImageHeight = int(float64(cam.ImageWidth) / cam.AspectRatio)

	// Additional camera settings
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * float64(cam.ImageWidth) / float64(cam.ImageHeight)

	// Calculate total width and height of viewport
	viewport_u := geometry.Vec3{viewport_width, 0, 0}
	viewport_v := geometry.Vec3{0, -viewport_height, 0}

	// Calculate displacement size for horizontal and vertical vectors
	cam.Pixel_delta_u = viewport_u.DivS(float64(cam.ImageWidth))
	cam.Pixel_delta_v = viewport_v.DivS(float64(cam.ImageHeight))

	// Caluclate position of top left corner
	viewport_upper_left := cam.Center.Sub(geometry.Vec3{0, 0, focal_length})
	viewport_upper_left = viewport_upper_left.Sub(viewport_u.DivS(2))
	viewport_upper_left = viewport_upper_left.Sub(viewport_v.DivS(2))

	// Calculate center of upper left pixel
	cam.Pixel_00_loc = viewport_upper_left
	cam.Pixel_00_loc = cam.Pixel_00_loc.Add(cam.Pixel_delta_u.MulS(0.5))
	cam.Pixel_00_loc = cam.Pixel_00_loc.Add(cam.Pixel_delta_v.MulS(0.5))
}


func (cam *Camera) Render(world *geometry.World) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, cam.ImageWidth, cam.ImageWidth))
	bar := progressbar.NewOptions(cam.ImageWidth * cam.ImageHeight)

	for i := 0; i < cam.ImageWidth; i++ {
		for j := 0; j < cam.ImageHeight; j++ {

			// Calculate next pixel center
			pixel_center := cam.Pixel_00_loc.Add(cam.Pixel_delta_u.MulS(float64(i)))
			pixel_center = pixel_center.Add(cam.Pixel_delta_v.MulS(float64(j)))

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
	return img
}


func (cam *Camera) RayColor(ray geometry.Ray, world *geometry.World) color.RGBA {
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

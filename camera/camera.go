package camera

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"raytracer/geometry"
	"github.com/schollz/progressbar/v3"
)

const MIN_DIST = 0.001
const MAX_DIST = 10000

type Camera struct {
	Center geometry.Vec3
	AspectRatio float64
	ImageWidth int
	ImageHeight int
	Pixel_00_loc geometry.Vec3
	Pixel_delta_u geometry.Vec3
	Pixel_delta_v geometry.Vec3
	SamplesPerPixel int
	MaxBounces int
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
	interval := geometry.Interval{Min: 0, Max: 0.99999}

	for i := 0; i < cam.ImageWidth; i++ {
		for j := 0; j < cam.ImageHeight; j++ {

			// Calculate next pixel center
			pixel_center := cam.Pixel_00_loc.Add(cam.Pixel_delta_u.MulS(float64(i)))
			pixel_center = pixel_center.Add(cam.Pixel_delta_v.MulS(float64(j)))

			// Sample ray color
			c := geometry.Vec3{}
			for k := 0; k < cam.SamplesPerPixel; k++ {
				ray := cam.sampleRay(i, j)
				c = c.Add(cam.rayColor(ray, cam.MaxBounces, world))
			}

			// Normalize for number of samples per pixel 
			c = c.DivS(float64(cam.SamplesPerPixel))

			rgba := color.RGBA{
				uint8(255 * interval.Clamp(math.Sqrt(c[0]))),
				uint8(255 * interval.Clamp(math.Sqrt(c[1]))),
				uint8(255 * interval.Clamp(math.Sqrt(c[2]))),
				255,
			}

			img.Set(i, j, rgba)
			bar.Add(1)
		}
	}
	return img
}

func (cam *Camera) sampleRay(i, j int) geometry.Ray {
	pixel_center := cam.Pixel_00_loc
	pixel_center = pixel_center.Add(cam.Pixel_delta_u.MulS(float64(i)))
	pixel_center = pixel_center.Add(cam.Pixel_delta_v.MulS(float64(j)))

	pixel_sample := pixel_center.Add(cam.samplePixel())

	return geometry.Ray{Orig: cam.Center, Dir:pixel_sample.Sub(cam.Center)}
}

func (cam *Camera) samplePixel() geometry.Vec3 {
	px := -0.5 + rand.Float64()
	py := -0.5 + rand.Float64()
	return cam.Pixel_delta_u.MulS(px).Add(cam.Pixel_delta_v.MulS(py))
}

func (cam *Camera) rayColor(ray geometry.Ray, depth int, world *geometry.World) geometry.Vec3 {
	record := geometry.HitRecord{}

	if depth < 1 {
		return geometry.Vec3{}
	}

	if world.Hit(ray, geometry.Interval{Min: MIN_DIST, Max: MAX_DIST}, &record) {
		ray_scattered := geometry.Ray{}
		attenuation := geometry.Vec3{}

		if record.Material.Scatter(&ray, &record, &attenuation, &ray_scattered) {
			return attenuation.Mul(cam.rayColor(ray_scattered, depth - 1, world))
		}
	}

	unit_ray := ray.Dir.Normalize()
	a := 0.5 * (unit_ray[1] + 1.0)
	c := geometry.Vec3{1.0, 1.0, 1.0}.MulS(1.0 - a)
	c = c.Add(geometry.Vec3{0.5, 0.7, 1.0}.MulS(a))

	return geometry.Vec3{c[0], c[1], c[2]}
}

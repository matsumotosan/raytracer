package camera

import (
	"image/color"

	"raytracer/geometry"
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


// func (camera *Camera) Initialize() {
// 	
// }


// func (camera Camera) Render(world geometry.Hittable) *image.RGBA {
// 	
// 	return img
// }


func (camera Camera) RayColor(ray geometry.Ray, world geometry.Hittable) color.RGBA {
	record := geometry.HitRecord{}

	if world.Hit(ray, geometry.Interval{0, 999999}, &record) {
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

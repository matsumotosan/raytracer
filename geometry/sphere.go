package geometry

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func (s Sphere) Hit(ray Ray, ray_t Interval, record *HitRecord) bool {
	oc := ray.Orig.Sub(s.Center)
	a := ray.Dir.Dot(ray.Dir)
	half_b := oc.Dot(ray.Dir)
	c := oc.Dot(oc) - s.Radius * s.Radius

	disc := half_b * half_b - a * c
	if disc < 0.0 {
		return false
	}

	sqrt_disc := math.Sqrt(disc)
	root := (-half_b - sqrt_disc) / a
	if !ray_t.surrounds(root) {
		root = (-half_b + sqrt_disc) / a
		if !ray_t.surrounds(root) {
			return false
		}
	}

	record.T = root
	record.Point = ray.At(record.T)
	outward_normal := (record.Point.Sub(s.Center)).DivS(s.Radius)
	record.SetFaceNormal(ray, outward_normal)

	return true
}

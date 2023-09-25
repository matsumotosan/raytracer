package geometry

import (
	"math"
	"math/rand"
)

const TOL = 1e-8

type Material interface {
	Scatter(ray_in *Ray, record *HitRecord, attenuation *Vec3, ray_scattered *Ray) bool 
}

type Lambertian struct {
	Albedo Vec3
}

type Metal struct {
	Albedo Vec3
	Fuzz float64
}

type Dielectric struct {
	RefractiveIndex float64
}

func (lam *Lambertian) Scatter(ray_in *Ray, record *HitRecord, attenuation *Vec3, ray_scattered *Ray) bool {
	scatter_dir := record.Normal.Add(RandUnitVec())
	if scatter_dir.NearZero(TOL) {
		scatter_dir = record.Normal
	}

	*ray_scattered = Ray{record.Point, scatter_dir}
	*attenuation = lam.Albedo
	return true
}

func (metal *Metal) Scatter(ray_in *Ray, record *HitRecord, attenuation *Vec3, ray_scattered *Ray) bool {
	reflected_dir := Reflect(ray_in.Dir.Normalize(), record.Normal)
	*ray_scattered = Ray{record.Point, reflected_dir.Add(RandUnitVec().MulS(metal.Fuzz))}
	*attenuation = metal.Albedo
	return true
}

func (dielectric *Dielectric) Scatter(ray_in *Ray, record *HitRecord, attenuation *Vec3, ray_scattered *Ray) bool {
	*attenuation = Vec3{1.0, 1.0, 1.0}

	var refraction_ratio float64
	if record.FrontFace {
		refraction_ratio = 1.0 / dielectric.RefractiveIndex
	} else {
		refraction_ratio = dielectric.RefractiveIndex
	}

	unit_dir := ray_in.Dir.Normalize()
	cos_theta := math.Min(-unit_dir.Dot(record.Normal), 1.0)
	sin_theta := math.Sqrt(1.0 - cos_theta * cos_theta)

	cannot_refract := refraction_ratio * sin_theta > 1.0
	var dir Vec3

	if cannot_refract || (Reflectance(cos_theta, refraction_ratio) > rand.Float64()) {
		dir = Reflect(unit_dir, record.Normal)
	} else {
		dir = Refract(unit_dir, record.Normal, refraction_ratio)
	}

	*ray_scattered = Ray{record.Point, dir}
	return true
}

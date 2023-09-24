package geometry

const TOL = 1e-8


type Material interface {
	Scatter(ray_in *Ray, record *HitRecord, attenuation *Vec3, ray_scattered *Ray) bool 
}


type Lambertian struct {
	Albedo Vec3
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


type Metal struct {
	Albedo Vec3
}

func (metal *Metal) Scatter(ray_in *Ray, record *HitRecord, attenuation *Vec3, ray_scattered *Ray) bool {
	reflected_dir := Reflect(ray_in.Dir.Normalize(), record.Normal)
	*ray_scattered = Ray{record.Point, reflected_dir}
	*attenuation = metal.Albedo
	return true
}

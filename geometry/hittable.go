package geometry


type HitRecord struct {
	Point Vec3
	Normal Vec3
	T float64
	FrontFace bool
	Material Material
}

func (hr *HitRecord) SetFaceNormal(ray Ray, outward_normal Vec3) {
	hr.FrontFace = ray.Dir.Dot(outward_normal) < 0.0
	if hr.FrontFace {
		hr.Normal = outward_normal
	} else {
		hr.Normal = outward_normal.MulS(-1)
	}
}


type Object interface {
	Hit(ray Ray, interval Interval, record *HitRecord) bool
}

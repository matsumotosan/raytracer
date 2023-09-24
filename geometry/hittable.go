package geometry


type HitRecord struct {
	Point Vec3
	Normal Vec3
	T float64
	FrontFace bool
}

func (hr *HitRecord) SetFaceNormal(ray Ray, outward_normal Vec3) {
	hr.FrontFace = ray.Dir.Dot(outward_normal) < 0.0
	if hr.FrontFace {
		hr.Normal = outward_normal
	} else {
		hr.Normal = outward_normal.MulS(-1)
	}
}

type Object interface { Hit(ray Ray, interval Interval, record *HitRecord) bool }

type World struct { Objects []Object }

func (world *World) Clear() { world.Objects = []Object{} }

func (world *World) Add(h Object) { world.Objects = append(world.Objects, h) }

func (world *World) Hit(ray Ray, ray_t Interval, record *HitRecord) bool {
	temp_rec := HitRecord{}
	hit := false
	closest := ray_t.Max

	for _, object := range world.Objects {
		if object.Hit(ray, Interval{ray_t.Min, closest}, &temp_rec) {
			hit = true
			closest = temp_rec.T
			*record = temp_rec
		}
	}

	return hit
}

package geometry

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

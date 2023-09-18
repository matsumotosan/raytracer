package geometry

type Ray struct {
	Orig Vec3
	Dir Vec3
}


func (r *Ray) At(t float64) Vec3 {
	return r.Orig.Add(r.Dir.MulS(t))
}

package geometry

import "math"

type Vec3 [3]float64


// vector-vector ops
func (u Vec3) Add(v Vec3) Vec3 {
	return Vec3{ u[0] + v[0], u[1] + v[1], u[2] + v[2] }
}

func (u Vec3) Sub(v Vec3) Vec3 {
	return Vec3{ u[0] - v[0], u[1] - v[1], u[2] - v[2] }
}

func (u Vec3) Mul(v Vec3) Vec3 {
	return Vec3{ u[0] * v[0], u[1] * v[1], u[2] * v[2] }
}

func (u Vec3) Div(v Vec3) Vec3 {
	return Vec3{ u[0] / v[0], u[1] / v[1], u[2] / v[2] }
}

func (u Vec3) Dot(v Vec3) float64 {
	return u[0] * v[0] + u[1] * v[1] + u[2] * v[2]
}

func (u Vec3) Cross(v Vec3) Vec3 {
	return Vec3{ u[0] + v[0], u[1] + v[1], u[2] + v[2] }
}


// vector-scalar ops
func (u Vec3) AddS(s float64) Vec3 {
	return Vec3{ u[0] + s, u[1] + s, u[2] + s }
}

func (u Vec3) SubS(s float64) Vec3 {
	return Vec3{ u[0] - s, u[1] - s, u[2] - s }
}

func (u Vec3) MulS(s float64) Vec3 {
	return Vec3{ u[0] * s, u[1] * s, u[2] * s }
}

func (u Vec3) DivS(s float64) Vec3 {
	return Vec3{ u[0] / s, u[1] / s, u[2] / s }
}



// vector norm
func (u Vec3) Norm() float64 {
	return math.Sqrt(u.Dot(u))
}


// normalize vector
func (u Vec3) GetUnit() Vec3 {
	return u.DivS(u.Norm())
}

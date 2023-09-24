package geometry

import (
	"math"
	"math/rand"
)

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


// random vectors
func RandVec(min, max float64) Vec3 {
	return Vec3{
		min + (max - min) * rand.Float64(),
		min + (max - min) * rand.Float64(),
		min + (max - min) * rand.Float64(),
	}
}

func RandVecInUnitSphere() Vec3 {
	for {
		v := RandVec(-1, 1)
		if v.Norm() < 1 {
			return v
		}
	}
}

func RandUnitVec() Vec3 {
	return RandVecInUnitSphere().GetUnit()
}

func RandVecOnHemisphere(normal Vec3) Vec3 {
	on_unit_sphere := RandUnitVec()
	if normal.Dot(on_unit_sphere) > 0.0 {
		return on_unit_sphere
	} else {
		return on_unit_sphere.MulS(-1)
	}
}

package src

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Unit() Vector {
	// TODO Guard against 0 magnitude
	return v.Multiply(1.0 / v.Magnitude())
}

func (v Vector) Add(a Vector) Vector {
	return Vector{
		X: v.X + a.X,
		Y: v.Y + a.Y,
	}
}

func (v Vector) Subtract(a Vector) Vector {
	return Vector{
		X: v.X - a.X,
		Y: v.Y - a.Y,
	}
}

func (v Vector) Multiply(m float64) Vector {
	return Vector{
		X: m * v.X,
		Y: m * v.Y,
	}
}

func Distance(a Vector, b Vector) float64 {
	return b.Subtract(a).Magnitude()
}

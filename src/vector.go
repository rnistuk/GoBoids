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
	n := v.Magnitude()
	return Vector{
		X: v.X / n,
		Y: v.Y / n,
	}
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
	return b.Add(a.Multiply(-1.0)).Magnitude()
}

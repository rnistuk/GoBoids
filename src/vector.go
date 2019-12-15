package src

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (b Vector) DirectionRadians() float64 {
	return math.Atan2(b.Y, b.X)
}

func (b Vector) DirectionDegrees() float64 {
	return b.DirectionRadians() * 180 / math.Pi
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

func (v Vector) ToString() string {
	return fmt.Sprintf("%f, %f", v.X, v.Y)
}

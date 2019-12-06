package src

import "github.com/veandco/go-sdl2/sdl"

type Boid struct {
	Position Vector
	Velocity Vector
}

// TODO: boid should not need to know about screen width or height, I need to refactor Boid to be only logic
func (b Boid) Draw(render *sdl.Renderer, screenWidth int32, screenHeight int32) {
	err := render.SetDrawColor(5, 55, 8, 255)
	if err != nil {
		return
	}
	r := sdl.Rect{screenWidth/2 + int32(b.Position.X/10.0), screenHeight/2 + int32(b.Position.Y/10.0), 5, 5}
	_ = render.FillRect(&r)
}

func NewBoid() Boid {
	return Boid{
		Position: Vector{},
		Velocity: Vector{},
	}
}

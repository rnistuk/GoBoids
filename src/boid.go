package src

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

func randomPosition(w float64) float64 {
	// TODO: check that w is greater than zero
	return float64(rand.Uint32()%uint32(w)) - float64(w/2.0)
}

func randomSpeed(width uint32) float64 {
	return float64(rand.Uint32()%width) - float64(width/2.0)
}

func NewBoid(renderer *sdl.Renderer, w *World) *Element {
	boid := &Element{}
	boid.Position = Vector{X: randomPosition(w.Size.X / 2.0), Y: randomPosition(w.Size.Y / 2.0)}
	boid.Velocity = Vector{X: 0, Y: 0}
	boid.Active = true
	boid.Rotation = boid.Velocity.DirectionDegrees()
	boid.World = w

	c := make([]Component, 0)
	boid.Components = &c

	//boid.addComponent(NewSpriteRenderer(boid, renderer, "resources/boid.bmp"))
	boid.addComponent(NewCohesionRule(boid))
	return boid
}

/*
func NewBoid(render *sdl.Renderer) *element {
	boid := &element{}
	boid.position = Vector{
		X: float64(rand.Uint32() % 500),
		Y: float64(rand.Uint32() % 500),
	}
	boid.active = true
	return boid
}*/
/*
// TODO: boid should not need to know about screen width or height, I need to refactor Boid to be only logic
func (b Boid) Draw(render *sdl.Renderer, screenWidth int32, screenHeight int32) {
	err := render.SetDrawColor(5, 55, 8, 255)
	if err != nil {
		return
	}
	r := sdl.Rect{screenWidth/2 + int32(b.Position.X/10.0), screenHeight/2 + int32(b.Position.Y/10.0), 5, 5}
	_ = render.FillRect(&r)
}
*/
/*
func NewBoid() Boid {
	return Boid{
		Position: Vector{},
		Velocity: Vector{},
	}
}*/
/*
func CentreOfFlock(b []Boid) Vector {
	centre := Vector{}
	for i := range b {
		centre = centre.Add(b[i].Position)
	}
	return centre.Multiply(1.0 / float64(len(b)))
}
*/

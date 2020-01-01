package src

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Boid struct {
	Position       Vector
	Velocity       Vector
	targetDistance float64
}

func NewBoid() Boid {
	return Boid{
		Vector{},
		Vector{},
		0.0,
	}
}

// TODO: boid should not need to know about screen width or height, I need to refactor Boid to be only logic
func (b Boid) Draw(render *sdl.Renderer, screenWidth int32, screenHeight int32) {
	err := render.SetDrawColor(5, 55, 8, 255)
	if err != nil {
		return
	}
	r := sdl.Rect{screenWidth/2 + int32(b.Position.X/1.0), screenHeight/2 + int32(b.Position.Y/1.0), 5, 5}
	_ = render.FillRect(&r)
}

func (b Boid) toString() string {
	return fmt.Sprintf("P: %s    V: %s", b.Position.toString(), b.Velocity.toString())
}

func (b *Boid) UpdateVelocity(bs Boids) {
	b.Velocity = b.Velocity.Add(NCohesionRule(b, bs))
	b.Velocity = b.Velocity.Add(SeparationRule(b, bs))
	b.Velocity = b.Velocity.Add(AlignmentRule(b, bs))
	b.Velocity = b.Velocity.Add(HomeRule(b, bs))
	b.Velocity = b.Velocity.Add(LimitSpeedRule(b, bs))
	b.Velocity = b.Velocity.Add(MinimumSpeedRule(b, bs))
}

func (b *Boid) UpdatePosition() {
	b.Position = b.Position.Add(b.Velocity)
}

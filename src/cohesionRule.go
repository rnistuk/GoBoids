package src

import (
	"errors"
	"github.com/veandco/go-sdl2/sdl"
)

type CohesionRule struct {
	boid *Element
}

func NewCohesionRule(b *Element) *CohesionRule {
	return &CohesionRule{b}
}

func (c *CohesionRule) OnUpdate() error {
	if c.boid.World.Elements == nil {
		return errors.New("No boids in this world")
	}

	cm := Vector{0.0, 0.0}
	for _, b := range *c.boid.World.Elements {
		cm = cm.Add(b.Position)
	}
	m := 1.0 / float64(len(*c.boid.World.Elements))
	cm = cm.Multiply(m)
	dv := c.boid.Position.Subtract(cm).Multiply(0.1)
	c.boid.Velocity = c.boid.Velocity.Add(dv)
	return nil
}

func (CohesionRule) OnDraw(_ *sdl.Renderer) error {
	return nil
}

package src

import (
	"testing"
)

func TestRule01(t *testing.T) {
	var boids []Boid
	boids = append(boids, NewBoid())
	boids = append(boids, NewBoid())

	boids[1].Position.X = 1.0

	v := Rule01(1, boids)
	if v.X+1.0 > smallFloat {
		t.Error(`Two boids did not fly to com`, v.X)
	}

}

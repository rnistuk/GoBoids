package src

import (
	"fmt"
	"testing"
)

func TestRule01(t *testing.T) {
	var boids []Boid
	boids = append(boids, NewBoid())
	boids = append(boids, NewBoid())

	boids[1].Position.X = 1.0

	velocity := Rule01(1, boids)

	fmt.Println(velocity.X)
	fmt.Println(velocity.Y)

}

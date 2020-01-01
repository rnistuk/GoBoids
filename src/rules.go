package src

import "fmt"

func NCohesionRule(b *Boid, bs Boids) Vector {
	// fly to the center of mass
	return CentreOfFlock(bs).Subtract(b.Position).Multiply(0.05)
}

func SeparationRule(b *Boid, bs Boids) Vector {
	var c Vector
	for _, ob := range bs {
		if ob != *b {
			d := Distance(b.Position, ob.Position)
			if d < 10 {
				c = c.Subtract(b.Position.Subtract(ob.Position)).Multiply(2.0)
			}
		}
	}
	return c
}

func AlignmentRule(b *Boid, bs Boids) Vector {
	// Boids try to match velocity with near boids.
	var pvj Vector
	for i := range bs {
		if &bs[i] != b {
			pvj = pvj.Add(bs[i].Velocity)
		}
	}
	pvj = pvj.Multiply(1.0 / float64(len(bs)-1))

	return pvj.Subtract(b.Velocity).Multiply(0.01)
}

func HomeRule(b *Boid, _ Boids) Vector {
	// Tendency towards a particular place
	place := Vector{0, 0}
	return place.Subtract(b.Position.Multiply(0.0050))
}

func LimitSpeedRule(b *Boid, _ Boids) Vector {
	// Limiting the speed
	maxSpeed := 20.0
	currentSpeed := b.Velocity.Magnitude()
	if currentSpeed > maxSpeed {
		return b.Velocity.Unit().Multiply(0.01 * (maxSpeed - currentSpeed))
	}
	return Vector{0.0, 0.0}
}

func MinimumSpeedRule(b *Boid, _ Boids) Vector {
	minSpeed := 5.0
	currentSpeed := b.Velocity.Magnitude()

	if currentSpeed < minSpeed {
		ds := minSpeed - currentSpeed
		fmt.Printf("ds: %f\n", ds)
		return b.Velocity.Unit().Multiply(1.0 * ds)
	}
	return Vector{0.0, 0.0}
}

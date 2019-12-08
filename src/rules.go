package src

func Rule01(j int, b []Boid) Vector {
	// fly to the center of mass
	return CentreOfFlock(b).Subtract(b[j].Position).Multiply(1.0 / 100.0)
}

func Rule02(j int, b []Boid) Vector {
	// keep a minimum distance between boids
	var c Vector
	for i := range b {
		if i != j {
			if 300 > Distance(b[j].Position, b[i].Position) {
				c = c.Subtract(b[i].Position.Subtract(b[j].Position).Multiply(0.5))
			}
		}
	}
	return c
}

func Rule03(j int, b []Boid) Vector {
	// Boids try to match velocity with near boids.
	var pvj Vector
	for i := range b {
		if i != j {
			pvj = pvj.Add(b[i].Velocity)
		}
	}
	pvj = pvj.Multiply(1.0 / float64(len(b)-1))

	return pvj.Subtract(b[j].Velocity).Multiply(0.25)
}

func Rule04(j int, b []Boid) Vector {
	// Tendency towards a particular place
	place := Vector{0, 0}
	return place.Subtract(b[j].Position.Multiply(1.0 / 10))
}

func Rule05(j int, b []Boid) Vector {
	// Limiting the speed
	s_max := 260.0
	s_current := b[j].Velocity.Magnitude()
	if s_current > s_max {
		return b[j].Velocity.Unit().Multiply(s_max - s_current)
	}
	return Vector{0.0, 0.0}
}

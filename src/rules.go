package src

var Parameters = map[string]float64{
	"near":       20.0,
	"cohesion":   0.005,
	"separation": 1.0,
	"alignment":  0.01,
	"home":       0.001,
	"minSpeed":   10.0,
	"maxSpeed":   20.0,
}

func NCohesionRule(b *Boid, bs Boids) Vector {
	fbs := FilterSortedByDistance(bs, Parameters["near"])
	return CentreOfFlock(fbs).Subtract(b.Position).Multiply(Parameters["cohesion"])
}

func SeparationRule(b *Boid, bs Boids) Vector {
	var c Vector
	for _, ob := range FilterSortedByDistance(bs, Parameters["near"]) {
		if ob != *b {
			d := Distance(b.Position, ob.Position)
			if d < 10 {
				c = c.Subtract(b.Position.Subtract(ob.Position)).Multiply(Parameters["separation"])
			}
		}
	}
	return c
}

func AlignmentRule(b *Boid, bs Boids) Vector {
	// Boids try to match velocity with near boids.
	var pvj Vector

	for i := range FilterSortedByDistance(bs, Parameters["near"]) {
		if &bs[i] != b {
			pvj = pvj.Add(bs[i].Velocity)
		}
	}
	pvj = pvj.Multiply(1.0 / float64(len(bs)-1))

	return pvj.Subtract(b.Velocity).Multiply(Parameters["alignment"])
}

func HomeRule(b *Boid, _ Boids) Vector {
	// Tendency towards a particular place
	place := Vector{0, 0}
	return place.Subtract(b.Position.Multiply(Parameters["home"]))
}

func LimitSpeedRule(b *Boid, _ Boids) Vector {
	// Limiting the speed
	maxSpeed := Parameters["maxSpeed"]
	currentSpeed := b.Velocity.Magnitude()
	if currentSpeed > maxSpeed {
		return b.Velocity.Unit().Multiply(0.01 * (maxSpeed - currentSpeed))
	}
	return Vector{0.0, 0.0}
}

func MinimumSpeedRule(b *Boid, _ Boids) Vector {
	minSpeed := Parameters["minSpeed"]
	currentSpeed := b.Velocity.Magnitude()

	if currentSpeed < minSpeed {
		ds := minSpeed - currentSpeed
		return b.Velocity.Unit().Multiply(1.0 * ds)
	}
	return Vector{0.0, 0.0}
}

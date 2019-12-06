package src

const NBoids = 100

func Rule01(j int, b [NBoids]Boid) Vector {
	var pos Vector
	for i := range b {
		if i != j {
			pos.X = pos.X + b[i].Position.X
			pos.Y = pos.Y + b[i].Position.Y
		}
	}
	pos.X = pos.X / 9
	pos.Y = pos.Y / 9

	var v Vector
	v.X = (pos.X - b[j].Position.X) / 100.0
	v.Y = (pos.Y - b[j].Position.Y) / 100.0
	return Vector{}
}

func Rule02(j int, b [NBoids]Boid) Vector {
	var c Vector
	for i := range b {
		if i != j {
			if 200 > Distance(b[j].Position, b[i].Position) {
				c.X = c.X - (b[i].Position.X-b[j].Position.X)*0.5
				c.Y = c.Y - (b[i].Position.Y-b[j].Position.Y)*0.5
			}
		}
	}
	return c
}

func Rule03(j int, b [NBoids]Boid) Vector {
	var pvj Vector
	for i := range b {
		if i != j {
			pvj.X += b[i].Velocity.X
			pvj.Y += b[i].Velocity.Y
		}
	}

	pvj.X /= NBoids - 1
	pvj.Y /= NBoids - 1

	pvj.X = (pvj.X - b[j].Velocity.X) / 4
	pvj.Y = (pvj.Y - b[j].Velocity.Y) / 4
	return pvj
}

func Rule04(j int, b [NBoids]Boid) Vector {
	place := Vector{300, 300}
	return Vector{(place.X - b[j].Position.X) / 100.0, (place.Y - b[j].Position.Y) / 100.0}
}

func Rule05(j int, b [NBoids]Boid) Vector {
	s_max := 260.0
	s_current := b[j].Velocity.Magnitude()

	if s_current > s_max {
		ds := s_max - s_current
		vc := b[j].Velocity.Unit()
		return vc.Multiply(ds)
	}
	return Vector{0.0, 0.0}
}

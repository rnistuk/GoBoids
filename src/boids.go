package src

type Boids []Boid
type PBoids *Boid
type BoidsP []*Boid

func CentreOfFlock(b Boids) Vector {
	centre := Vector{}
	for i := range b {
		centre = centre.Add(b[i].Position)
	}
	return centre.Multiply(1.0 / float64(len(b)))
}

func setTargetDistances(b Boid, bs Boids) Boids {
	for i := range bs {
		bs[i].targetDistance = Distance(b.Position, bs[i].Position)
	}
	return bs
}

type ByDistance Boids

func (bs ByDistance) Len() int           { return len(bs) }
func (bs ByDistance) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs ByDistance) Less(i, j int) bool { return bs[i].targetDistance < bs[j].targetDistance }

func FilterSortedByDistance(b Boid, bs Boids, d float64) Boids {
	var fbs Boids
	bs = setTargetDistances(b, bs)
	for _, b := range bs {
		if b.targetDistance <= d {
			fbs = append(fbs, b)
		}
	}
	return fbs
}

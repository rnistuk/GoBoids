package src

import (
	"fmt"
	"sort"
)

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
	for _, ob := range bs {
		ob.targetDistance = Distance(b.Position, ob.Position)
	}
	return bs
}

type ByDistance Boids

func (bs ByDistance) Len() int           { return len(bs) }
func (bs ByDistance) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs ByDistance) Less(i, j int) bool { return bs[i].targetDistance < bs[j].targetDistance }

func SortClosest(b Boid, bs Boids) Boids {
	bs = setTargetDistances(b, bs)
	sort.Sort(ByDistance(bs))
	return bs
}

func FilterSortedByDistance(bs Boids, d float64) Boids {
	var fbs Boids
	for _, b := range bs {
		if b.targetDistance <= d {
			fbs = append(fbs, b)
		}
	}
	return fbs
}

func stats(boids Boids) string {
	var min, max, ave float64

	min = boids[0].Velocity.Magnitude()
	for _, b := range boids {
		s := b.Velocity.Magnitude()
		if min > s {
			min = s
		}
		if max < s {
			max = s
		}
		ave = ave + s
	}
	ave = ave / float64(len(boids))
	return fmt.Sprintf("min: %f ave: %f  max: %f", min, ave, max)
}

package src

/*
func TestRule01(t *testing.T) {
	var boids []Boid
	angle := 0.0
	for i := 0; i < 3; i++ {
		boids = append(boids, NewBoid())
		boids[i].Position = Vector{10.0 * math.Cos(angle), 10.0 * math.Sin(angle)}
		fmt.Println(boids[i].Position.toString())
		angle += 2.0 * math.Pi / 3.0
	}

	v := Rule01(0, boids)
	if v.X+0.1 > smallFloat {
		t.Error(`Rule did not calculate the velocity to fly to the C of M, v.X:`, v.toString())
	}

	v = Rule01(1, boids)
	fmt.Print(v.toString())

	vAccepted := Vector{5.000000, -8.660254}.Multiply(1.0 / 100.0)
	difference := vAccepted.Subtract(v).Magnitude()

	if difference > smallFloat {
		t.Error(`Rule did not calculate the velocity to fly to the C of M, v.X:`, v.toString())
	}
}
*/

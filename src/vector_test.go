package src

import (
	"math"
	"testing"
)

const smallFloat = 0.00001

func TestVector_Add(t *testing.T) {
	suta := Vector{0.0, 0.0}
	sutb := Vector{0.0, 0.0}
	accepted := Vector{0.0, 0.0}
	if suta.Add(sutb) != accepted {
		t.Error(`Add test 1 is false`)
	}

	sutb.X = -1.0
	accepted.X = -1.0
	if suta.Add(sutb) != accepted {
		t.Error(`Add test 2 is false`)
	}

	sutb.X = 0.0
	sutb.Y = -1.0
	accepted.X = 0.0
	accepted.Y = -1.0
	if suta.Add(sutb) != accepted {
		t.Error(`Add test 3 is false`)
	}

}

func TestVector_Magnitude(t *testing.T) {
	sut := Vector{
		X: 0,
		Y: 0,
	}

	if sut.Magnitude() > smallFloat {
		t.Error(`sut{0.0}.Magnitude() == 0.0 is false`)
	}

	sut.X = 1.0
	if math.Abs(sut.Magnitude()-1.0) > smallFloat {
		t.Error(`sut{1.0,0}.Magnitude() == 1.0 is false`)
	}

	sut.X = 0.0
	sut.Y = 1.0
	if math.Abs(sut.Magnitude()-1.0) > smallFloat {
		t.Error(`sut{1.0,0}.Magnitude() == 1.0 is false`)
	}

	sut.X = 3.0
	sut.Y = 4.0
	if math.Abs(sut.Magnitude()-5.0) > smallFloat {
		t.Error(`sut{3.0,4.0}.Magnitude() == 5.0 is false`)
	}

}

func TestVector_Unit(t *testing.T) {
	var sut = Vector{3.0, 4.0}
	sutUnit := sut.Unit()
	var accepted = Vector{X: 3.0 / 5.0, Y: 4.0 / 5.0}
	if math.Abs(sutUnit.Magnitude()-accepted.Magnitude()) > smallFloat {
		t.Error(`sut.Unit() == accepted is false`)
	}
}

func TestVector_Multiply(t *testing.T) {
	sut := Vector{0.0, 0.0}.Multiply(5.0)
	if math.Abs(sut.Magnitude()) > smallFloat {
		t.Error(`Magnitude of zero vector multiplied by 5.0 must be zero`)
	}

	sut = Vector{37.4525, 79.32487392}.Unit()
	sut = sut.Multiply(5.0)
	if sut.Magnitude()-5 > smallFloat {
		t.Error(`Magnitude of unit vector multiplied by 5.0 must be 5.0`)
	}
}

func TestVector_Subtract(t *testing.T) {
	sutA := Vector{0.0, 0.0}
	sutB := Vector{0.0, 0.0}

	if sutA.Subtract(sutB).Magnitude() > smallFloat {
		t.Error(`Magnitude of zero vector subtracted from zero vector must be zero`)
	}

	sutB = Vector{3.0, 4.0}
	acc := sutA.Subtract(sutB)
	if math.Abs(acc.X+3.0) > smallFloat {
		t.Error(`X param of sut incorrect after subtraction`)
	}

	if math.Abs(acc.Y+4.0) > smallFloat {
		t.Error(`Y param of sut incorrect after subtraction`)
	}

	sutA = Vector{-3.0, -4.0}
	acc = sutA.Subtract(sutB)

	acc = sutA.Subtract(sutB)
	if math.Abs(acc.X+6.0) > smallFloat {
		t.Error(`X param of sut incorrect after subtraction`)
	}

	if math.Abs(acc.Y+8.0) > smallFloat {
		t.Error(`Y param of sut incorrect after subtraction`)
	}

}

func TestDistance(t *testing.T) {
	sutA := Vector{0, 0}

	if Distance(sutA, sutA) > smallFloat {
		t.Error(`Distance between a vector and itself must be zero`)
	}

	sutB := Vector{5.0, 0.0}
	if Distance(sutA, sutB)-5.0 > smallFloat {
		t.Error(`Distance between a vector and a {5,0} vector must be 5.0`)
	}

	sutB = Vector{-5.0, 0.0}
	if Distance(sutA, sutB)-5.0 > smallFloat {
		t.Error(`Distance between a vector and a {-5,0} vector must be 5.0`)
	}

	sutB = Vector{-3.0, 4.0}
	if Distance(sutA, sutB)-5.0 > smallFloat {
		t.Error(`Distance between a zero vector and a {-3,4} vector must be 5.0`)
	}

}

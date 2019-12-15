package src

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type Element struct {
	Position   Vector
	Velocity   Vector
	Rotation   float64 // degrees (?) sdl likes things in degrees
	Active     bool
	Components *[]Component
	World      *World
}

func (e *Element) Draw(renderer *sdl.Renderer) error {
	for _, comp := range *e.Components {
		err := comp.OnDraw(renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Element) Update() error {
	for _, comp := range *e.Components {
		err := comp.OnUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Element) addComponent(newComponent Component) {
	for _, component := range *e.Components {
		if reflect.TypeOf(newComponent) == reflect.TypeOf(component) {
			panic(fmt.Sprintf("component of type %v already exists in element", reflect.TypeOf(newComponent)))
		}
	}
	*e.Components = append(*e.Components, newComponent)
}

func (e *Element) getComponent(c Component) Component {
	t := reflect.TypeOf(c)
	for _, cmp := range *e.Components {
		if reflect.TypeOf(cmp) == t {
			return c
		}
	}
	panic(fmt.Sprintf("component of type %v does not exist in element", t))
}

func (b Element) ToString() string {
	return fmt.Sprintf("P: %s    V: %s", b.Position.ToString(), b.Velocity.ToString())
}

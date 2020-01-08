package src

import (
	"github.com/veandco/go-sdl2/sdl"
)

type fn func()

type EventMapType map[sdl.Rect]fn

func (e *EventMapType) FireEvent(pt sdl.Point) {
	for r, f := range *e {
		if pt.InRect(&r) {
			f()
			return
		}
	}
	return
}

func (e *EventMapType) AddButton(r sdl.Rect, f fn) {
	(*e)[r] = f
}

package src

import "github.com/veandco/go-sdl2/sdl"

type Component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
}

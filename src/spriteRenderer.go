package src

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container     *Element
	tex           *sdl.Texture
	width, height float64
}

func NewSpriteRenderer(container *Element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	tex := textureFromBMP(renderer, filename)
	_, _, width, height, err := tex.Query()

	if err != nil {
		panic(fmt.Errorf("Querying texture %v", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       tex,
		width:     float64(width),
		height:    float64(height),
	}
}

func (sr *spriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	x := sr.container.Position.X - sr.width/2.0
	y := sr.container.Position.Y - sr.height/2.0

	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(sr.width), H: int32(sr.height)},
		&sdl.Rect{X: int32(x + 400), Y: int32(y + 400), W: int32(sr.width / 2), H: int32(sr.height / 2)},
		sr.container.Rotation,
		&sdl.Point{X: int32(sr.width) / 2, Y: int32(sr.height) / 2},
		sdl.FLIP_NONE)

	return nil
}

func (sr *spriteRenderer) OnUpdate() error {
	return nil
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	defer img.Free()
	if err != nil {
		panic(fmt.Errorf("Could not load sprite named: %v, %v", filename, sdl.GetError()))
	}
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("Could not create sprite texture named: %v", filename))
	}
	return tex
}

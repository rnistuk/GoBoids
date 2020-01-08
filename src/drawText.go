package src

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func DrawTextInRect(r *sdl.Renderer, font *ttf.Font, txt string, rect sdl.Rect) {
	var err error
	var solidSurface *sdl.Surface
	var solidTexture *sdl.Texture

	if solidSurface, err = font.RenderUTF8Solid(txt, sdl.Color{R: 255, G: 255, A: 255}); err != nil {
		fmt.Printf("Failed to render text: %s\n", err)
	}

	if solidTexture, err = r.CreateTextureFromSurface(solidSurface); err != nil {
		fmt.Printf("Failed to create texture: %s\n", err)
	}

	solidSurface.Free()
	_ = r.Copy(solidTexture, nil, &rect)
	defer solidTexture.Destroy()
}

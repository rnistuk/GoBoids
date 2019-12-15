package main

import (
	"fmt"
	"github.com/goboids/src"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"time"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 800
)

func initSDL() (*sdl.Window, *sdl.Renderer) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing SDL:", err)
		panic(err)
	}
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, ScreenWidth, ScreenHeight, sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("Initializing window:", err)
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer:", err)
		panic(err)
	}

	return window, renderer
}

func main() {
	window, renderer := initSDL()
	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()

	rand.Seed(time.Now().UTC().UnixNano())

	var world = &src.World{
		OriginPoint: src.Vector{X: ScreenWidth, Y: ScreenHeight}.Multiply(0.5),
		Size:        src.Vector{ScreenWidth, ScreenHeight},
		Elements:    &[]src.Element{},
	}

	// Add boids
	nBoids := 2
	for i := 0; i < nBoids; i++ {
		*world.Elements = append(*world.Elements, *src.NewBoid(renderer, world))
	}

	running := true
	for running {
		Draw(renderer, *world.Elements)
		Update(world.Elements)

		handleEvents(&running)
		sdl.Delay(100)
	}
}

func Update(boids *[]src.Element) {
	for i := range *boids {
		_ = (*boids)[i].Update()

		fmt.Printf("velocity: %s\n", (*boids)[i].Velocity.ToString())
	}

	for _, b := range *boids {
		b.Position = b.Position.Add(b.Velocity)
	}

}

func Draw(renderer *sdl.Renderer, elements []src.Element) {
	drawWorld(renderer)
	for i := range elements {
		elements[i].Draw(renderer)
	}
	renderer.Present()
}

func drawWorld(renderer *sdl.Renderer) {
	renderer.SetDrawColor(147, 172, 207, 255)
	renderer.Clear()

	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawLine(ScreenHeight/2, 0, ScreenHeight/2, ScreenWidth)
	renderer.DrawLine(0, ScreenWidth/2, ScreenWidth, ScreenWidth/2)

}

func handleEvents(running *bool) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			*running = false
			break
		case *sdl.MouseButtonEvent:
			x, y, state := sdl.GetMouseState()
			fmt.Printf("event: %d %d %d\n", x, y, state)
			break

		}
	}
}

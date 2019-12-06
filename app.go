package main

import (
	"fmt"
	"github.com/goboids/src"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

const (
	NBoids       = src.NBoids
	ScreenWidth  = 800
	ScreenHeight = 800
)

func initializeBoids(b *[NBoids]src.Boid) {
	for i := range b {
		b[i].Position = src.Vector{
			X: float64(rand.Uint32() % ScreenWidth),
			Y: float64(rand.Uint32() % ScreenHeight),
		}
	}
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, ScreenWidth, ScreenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	var boids [NBoids]src.Boid

	initializeBoids(&boids)

	running := true
	for running {
		renderer.SetDrawColor(128, 255, 128, 255)
		renderer.Clear()

		draw_boids(renderer, boids)

		renderer.Present()

		update_boids(&boids)

		sdl.Delay(100)

		handleEvents(&running)
	}
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

func update_boids(boids *[NBoids]src.Boid) {
	for i := range boids {
		// TODO Learn how to implement the strategy pattern in Go
		v1 := src.Rule01(i, *boids)
		v2 := src.Rule02(i, *boids)
		v3 := src.Rule03(i, *boids)
		v4 := src.Rule04(i, *boids)
		v5 := src.Rule05(i, *boids)

		boids[i].Velocity.X += v1.X + v2.X + v3.X + v4.X + v5.X
		boids[i].Velocity.Y += v1.Y + v2.Y + v3.Y + v4.Y + v5.Y
		boids[i].Position.X += boids[i].Velocity.X
		boids[i].Position.Y += boids[i].Velocity.Y
	}
}

func draw_boids(renderer *sdl.Renderer, boids [NBoids]src.Boid) {
	for i := range boids {
		boids[i].Draw(renderer, ScreenWidth, ScreenHeight)
	}
}

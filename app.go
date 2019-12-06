package main

import (
	"fmt"
	"github.com/goboids/src"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 800
)

func initializeBoids(b *[]src.Boid, nBoids int) {
	for i := 0; i < nBoids; i++ {
		*b = append(*b, src.Boid{
			Position: src.Vector{
				X: float64(rand.Uint32() % ScreenWidth),
				Y: float64(rand.Uint32() % ScreenHeight),
			},
			Velocity: src.Vector{},
		})
	}
}

func main() {
	nBoids := 30
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

	var boids []src.Boid

	initializeBoids(&boids, nBoids)

	running := true
	for running {
		renderer.SetDrawColor(128, 255, 128, 255)
		renderer.Clear()

		draw_boids(renderer, boids)

		renderer.Present()

		boids = update_boids(boids)

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

func update_boids(boids []src.Boid) []src.Boid {
	var nextBoids []src.Boid
	nBoids := len(boids)
	for i := 0; i < nBoids; i++ {
		nextBoid := boids[i]
		// TODO Learn how to implement the strategy pattern in Go
		nextBoid.Velocity = nextBoid.Velocity.Add(src.Rule01(i, boids))
		nextBoid.Velocity = nextBoid.Velocity.Add(src.Rule02(i, boids))
		nextBoid.Velocity = nextBoid.Velocity.Add(src.Rule03(i, boids))
		nextBoid.Velocity = nextBoid.Velocity.Add(src.Rule04(i, boids))
		nextBoid.Velocity = nextBoid.Velocity.Add(src.Rule05(i, boids))

		nextBoid.Position = nextBoid.Position.Add(nextBoid.Velocity)
		nextBoids = append(nextBoids, nextBoid)
	}
	return nextBoids
}

func draw_boids(renderer *sdl.Renderer, boids []src.Boid) {
	for i := range boids {
		boids[i].Draw(renderer, ScreenWidth, ScreenHeight)
	}
}

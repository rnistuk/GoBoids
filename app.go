package main

import (
	"fmt"
	"github.com/goboids/src"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

const (
	ScreenWidth  = 1000
	ScreenHeight = 1000
)

func initializeBoids(b *[]src.Boid, nBoids int) {
	for i := 0; i < nBoids; i++ {
		*b = append(*b, src.Boid{
			Position: src.Vector{
				X: float64(rand.Uint32() % 50),
				Y: float64(rand.Uint32() % 50),
			},
			Velocity: src.Vector{},
		})
	}
}

func main() {
	nBoids := 7
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
		renderer.SetDrawColor(147, 172, 207, 255)
		renderer.Clear()

		draw_boids(renderer, boids)

		renderer.Present()

		update_boids(&boids)

		sdl.Delay(50)

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

func update_boids(boids *[]src.Boid) {
	nBoids := len(*boids)
	for i := 0; i < nBoids; i++ {
		(*boids)[i].UpdateVelocity(*boids)
	}

	for i := 0; i < nBoids; i++ {
		(*boids)[i].UpdatePosition()
	}
}

func draw_boids(renderer *sdl.Renderer, boids []src.Boid) {
	for i := range boids {
		boids[i].Draw(renderer, ScreenWidth, ScreenHeight)
	}
}

func stats(boids []src.Boid) string {
	var min, max, ave float64

	min = boids[0].Velocity.Magnitude()
	for _, b := range boids {
		s := b.Velocity.Magnitude()
		if min > s {
			min = s
		}
		if max < s {
			max = s
		}
		ave = ave + s
	}
	ave = ave / float64(len(boids))
	return fmt.Sprintf("min: %f ave: %f  max: %f", min, ave, max)
}

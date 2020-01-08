package main

import (
	"fmt"
	"github.com/goboids/src"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math/rand"
)

const (
	ScreenWidth  = 1000
	ScreenHeight = 1000
)

var (
	OpenUIButtonRect   = sdl.Rect{X: 0, Y: 0, W: 10, H: 10}
	OpenUIPaneRect     = sdl.Rect{X: 0, Y: 0, W: ScreenWidth / 5, H: ScreenHeight}
	DecRangeButtonRect = sdl.Rect{X: 5, Y: 20, W: 20, H: 20}
	IncRangeButtonRect = sdl.Rect{X: 95, Y: 20, W: 20, H: 20}
	UIIsOpen           = false
)

func initializeBoids(b *[]src.Boid, n int) {
	for i := 0; i < n; i++ {
		*b = append(*b, src.Boid{
			Position: src.Vector{
				X: float64(rand.Uint32() % 50),
				Y: float64(rand.Uint32() % 50),
			},
			Velocity: src.Vector{},
		})
	}
}

func Setup() (*sdl.Window, *sdl.Renderer) {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing SDL:", err)
	}

	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, ScreenWidth, ScreenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Initializing window:", err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer:", err)
	}

	if err = ttf.Init(); err != nil {
		fmt.Printf("Failed to initialize TTF: %s\n", err)
	}

	return window, renderer
}

func ShutDown(window *sdl.Window, renderer *sdl.Renderer) {
	sdl.Quit()
	_ = window.Destroy()
	_ = renderer.Destroy()
	ttf.Quit()
}

func main() {
	window, renderer := Setup()

	if window != nil && renderer != nil {
		var boids []src.Boid
		buttonEvents := src.EventMapType{}

		initializeBoids(&boids, 30)

		buttonEvents.AddButton(OpenUIButtonRect, func() { UIIsOpen = !UIIsOpen })
		buttonEvents.AddButton(IncRangeButtonRect, func() { src.Parameters["near"] = src.Parameters["near"] + 1 })
		buttonEvents.AddButton(DecRangeButtonRect, func() { src.Parameters["near"] = src.Parameters["near"] - 1 })

		running := true
		for running {
			drawBackground(renderer)

			drawUI(renderer)

			drawBoids(renderer, boids)

			renderer.Present()

			updateBoids(&boids)

			handleEvents(buttonEvents, &running)

			sdl.Delay(50)
		}
	}

	ShutDown(window, renderer)
}

func handleEvents(e src.EventMapType, running *bool) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quitting")
			*running = false
			break

		case *sdl.MouseButtonEvent:
			var pt sdl.Point
			var state uint32
			pt.X, pt.Y, state = sdl.GetMouseState()

			if state == 4 {
				e.FireEvent(pt)
				fmt.Printf("MouseButtonEvent: %d %d %d\n", pt.X, pt.Y, state)
			}
			break

		case *sdl.MouseWheelEvent:
			x, y, state := sdl.GetMouseState()
			fmt.Printf("MouseWheelEvent: %d %d %d\n", x, y, state)
			break
		}
	}
}

func drawBackground(r *sdl.Renderer) {
	_ = r.SetDrawColor(147, 172, 207, 255)
	_ = r.Clear()
}

func drawBoids(renderer *sdl.Renderer, boids []src.Boid) {
	for i := range boids {
		boids[i].Draw(renderer, ScreenWidth, ScreenHeight)
	}
}

func drawUI(r *sdl.Renderer) {
	_ = r.SetDrawColor(0, 0, 0, 255)

	_ = r.DrawRect(&OpenUIButtonRect)

	if UIIsOpen {
		var font *ttf.Font
		var err error
		if font, err = ttf.OpenFont("Verdana.ttf", 90); err != nil {
			fmt.Printf("Failed to open font: %s\n", err)
		}
		src.DrawTextInRect(r, font, fmt.Sprintf(" - Near + : %v", src.Parameters["near"]), sdl.Rect{W: 0, X: 190, H: 0, Y: 53})

		_ = r.SetDrawColor(0, 0, 0, 255)
		_ = r.DrawRect(&OpenUIPaneRect)
		_ = r.DrawRect(&IncRangeButtonRect)
		_ = r.DrawRect(&DecRangeButtonRect)
	}
}

func updateBoids(boids *[]src.Boid) {
	nBoids := len(*boids)
	for i := 0; i < nBoids; i++ {
		(*boids)[i].UpdateVelocity(*boids)
	}

	for i := 0; i < nBoids; i++ {
		(*boids)[i].UpdatePosition()
	}
}

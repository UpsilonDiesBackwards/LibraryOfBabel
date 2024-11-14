package main

import (
	"fmt"
	"github.com/UpsilonDiesBackwards/LibraryOfBabel/engine/io"
	"github.com/UpsilonDiesBackwards/LibraryOfBabel/engine/rendering"
	"github.com/UpsilonDiesBackwards/LibraryOfBabel/game"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"runtime"
	"time"
)

var DeltaTime float64

func main() {
	runtime.LockOSThread()
	defer glfw.Terminate()

	window, err := rendering.NewWindow("Library of Babel")
	if err != nil {
		fmt.Println(err)
	}

	app := game.NewApplication()
	app.Initialise()

	rend := rendering.NewRenderer(window)

	rend.NewObject("res/models/cube.obj", "", "char")

	rend.GetObject("char").SetPosition(mgl32.Vec3{0, -1, -4})
	rend.GetObject("char").SetScale(mgl32.Vec3{1, 1, 1})

	var previousTime = time.Now()
	for !window.ShouldClose() {
		DeltaTime = CalculateDeltaTime(previousTime)
		previousTime = time.Now()

		io.InputRunner(window, DeltaTime, app.Player.Camera, float64(app.Player.WalkSpeed))

		app.Run()

		rend.Draw(*app.Player.Camera)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func CalculateDeltaTime(previousTime time.Time) float64 {
	currentTime := time.Now()
	deltaTime := currentTime.Sub(previousTime).Seconds()
	return deltaTime
}

package io

import (
	"fmt"
	"github.com/UpsilonDiesBackwards/LibraryOfBabel/engine/rendering"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
)

var ViewportTransform mgl32.Mat4
var u = &UserInput{}

var isZooming = false

func InputRunner(win *rendering.Window, deltaTime float64, camera *rendering.Camera, zoomMult float64) error {
	var adjCSpeed float64

	if isZooming {
		adjCSpeed = deltaTime * float64(camera.Speed*zoomMult)
	} else {
		adjCSpeed = deltaTime * float64(camera.Speed)
	}

	if ActionState[VP_FORW] {
		horizontalFront := mgl64.Vec3{camera.Front.X(), 0, camera.Front.Z()}.Normalize()
		camera.Position = camera.Position.Add(horizontalFront.Mul(adjCSpeed))
	}
	if ActionState[VP_BACK] {
		horizontalFront := mgl64.Vec3{camera.Front.X(), 0, camera.Front.Z()}.Normalize()
		camera.Position = camera.Position.Sub(horizontalFront.Mul(adjCSpeed))
	}
	if ActionState[VP_LEFT] {
		camera.Position = camera.Position.Sub(camera.Front.Cross(camera.Up).Mul(adjCSpeed))
	}
	if ActionState[VP_RGHT] {
		camera.Position = camera.Position.Add(camera.Front.Cross(camera.Up).Mul(adjCSpeed))
	}

	if ActionState[VP_SPRINT] {
		isZooming = !isZooming
	}

	if ActionState[ED_QUIT] {
		fmt.Println("Exiting!")
		glfw.Terminate()
	}

	camera.UpdateDirection(u.CursorChange().X(), u.CursorChange().Y())
	u.CheckpointCursorChange()
	ViewportTransform = camera.GetTransform()
	InputManager(win, u)

	return nil
}

func GetUserInput() *UserInput {
	return u
}

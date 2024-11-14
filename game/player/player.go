package player

import (
	"github.com/UpsilonDiesBackwards/LibraryOfBabel/engine/rendering"
	"github.com/go-gl/mathgl/mgl64"
)

type Player struct {
	WalkSpeed       float64
	SprintSpeedMult float64

	Gravity float32

	Camera *rendering.Camera
}

func NewPlayer() *Player {
	return &Player{
		WalkSpeed:       6.0,
		SprintSpeedMult: 2, // value that the player's speed is multiplied by when sprinting.
		Gravity:         9.81,

		Camera: &rendering.Camera{
			Position:    mgl64.Vec3{0, 0, 0},
			Up:          mgl64.Vec3{0, 1, 0},
			WorldUp:     mgl64.Vec3{0, 1, 0},
			Yaw:         -90,
			Pitch:       0,
			Speed:       1,
			Sensitivity: 0.075,
			Fov:         60,
		},
	}
}

func (p *Player) Initialise() {
	p.Camera.Speed = p.WalkSpeed
}

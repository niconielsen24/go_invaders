package mobs

import (
	"go_invaders/game/bullet"
	"go_invaders/game/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const extraStepInterval = 0.35

type Extra struct {
	Sprite      rl.Texture2D
	BoundingBox internal.BoundingBox
	Position    internal.Point
	gameBounds  internal.BoundingBox
	stepTimer   float32
	alive       bool
}

func NewExtra(sprite rl.Texture2D, position internal.Point, gameBounds internal.BoundingBox) *Extra {
	return &Extra{
		Sprite:      sprite,
		BoundingBox: centeredBox(position, spriteSize),
		Position:    position,
		gameBounds:  gameBounds,
		alive:       true,
	}
}

// Mystery ship: fast horizontal sweep, disappears when it exits bounds.
func (e *Extra) Move() {
	e.stepTimer += rl.GetFrameTime()
	if e.stepTimer < extraStepInterval {
		return
	}
	e.stepTimer = 0
	e.Position.X += 20
	e.BoundingBox = centeredBox(e.Position, spriteSize)

	if e.Position.X > e.gameBounds.X+e.gameBounds.Width {
		e.alive = false
	}
}

func (e *Extra) Shoot() *bullet.Bullet { return nil }

func (e *Extra) CollidesWith(other internal.BoundingBox) bool {
	return e.BoundingBox.Intersects(other)
}

func (e *Extra) Draw() {
	rl.DrawTexture(e.Sprite, e.Position.X, e.Position.Y, rl.White)
}

func (e *Extra) IsAlive() bool { return e.alive }

func (e *Extra) Hit() {
	e.alive = false
}

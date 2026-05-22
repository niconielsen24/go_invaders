package mobs

import (
	"go_invaders/game/bullet"
	"go_invaders/game/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const yellowStepInterval = 0.8

type Yellow struct {
	Sprite      rl.Texture2D
	BoundingBox internal.BoundingBox
	Position    internal.Point
	stepTimer   float32
	direction   int32
	alive       bool
	hitPoints   int
}

func NewYellow(sprite rl.Texture2D, position internal.Point) *Yellow {
	return &Yellow{
		Sprite:      sprite,
		BoundingBox: centeredBox(position, spriteSize),
		Position:    position,
		direction:   1,
		alive:       true,
		hitPoints:   2,
	}
}

// Faster zigzag than Red, takes two hits to kill.
func (y *Yellow) Move() {
	y.stepTimer += rl.GetFrameTime()
	if y.stepTimer < yellowStepInterval {
		return
	}
	y.stepTimer = 0
	y.Position.X += 20 * y.direction
	y.Position.Y += 15
	y.direction *= -1
	y.BoundingBox = centeredBox(y.Position, spriteSize)
}

func (y *Yellow) Shoot() *bullet.Bullet { return nil }

func (y *Yellow) CollidesWith(other internal.BoundingBox) bool {
	return y.BoundingBox.Intersects(other)
}

func (y *Yellow) Draw() {
	rl.DrawTexture(y.Sprite, y.Position.X, y.Position.Y, rl.Yellow)
}

func (y *Yellow) IsAlive() bool { return y.alive }

func (y *Yellow) Hit() {
	y.hitPoints--
	if y.hitPoints <= 0 {
		y.alive = false
	}
}

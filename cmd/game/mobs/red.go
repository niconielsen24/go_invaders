package mobs

import (
	"go_invaders/game/bullet"
	"go_invaders/game/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const stepInterval = 1.5

type Red struct {
	Sprite      rl.Texture2D
	BoundingBox internal.BoundingBox
	Position    internal.Point
	stepTimer   float32
	direction   int32
	alive       bool
	hitPoints   int
}

const spriteSize = 50

func NewRed(sprite rl.Texture2D, position internal.Point) *Red {
	return &Red{
		Sprite:      sprite,
		BoundingBox: centeredBox(position, spriteSize),
		Position:    position,
		stepTimer:   0,
		direction:   1,
		alive:       true,
		hitPoints:   1,
	}
}

func (r *Red) Move() {
	r.stepTimer += rl.GetFrameTime()
	if r.stepTimer < stepInterval {
		return
	}
	r.stepTimer = 0
	r.Position.X += 20 * r.direction
	r.Position.Y += 20
	r.direction *= -1
	r.BoundingBox = centeredBox(r.Position, spriteSize)
}

func centeredBox(p internal.Point, size int32) internal.BoundingBox {
	half := size / 2
	return internal.NewBoundingBox(p.X-half, p.Y-half, size, size)
}

func (r *Red) CollidesWith(other internal.BoundingBox) bool {
	return r.BoundingBox.Intersects(other)
}

func (r *Red) Shoot() *bullet.Bullet { return nil }

func (r *Red) Draw() {
	rl.DrawTexture(r.Sprite, r.Position.X, r.Position.Y, rl.Red)
}

func (r *Red) IsAlive() bool {
	return r.alive
}

func (r *Red) Hit() {
	r.hitPoints--
	if r.hitPoints <= 0 {
		r.kill()
	}
}

func (r *Red) kill() {
	r.alive = false
}

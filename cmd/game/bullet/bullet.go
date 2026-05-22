package bullet

import (
	"go_invaders/game/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Up     = -1
	Down   = 1
	width  = 10
	height = 15
)

type Bullet struct {
	pos      internal.Point
	hitbox   internal.BoundingBox
	Velocity float32
	bounds   internal.BoundingBox
	collided bool
}

func NewBullet(x, y int32, velocity float32, boundedTo internal.BoundingBox) Bullet {
	return Bullet{
		pos:      internal.NewPoint(x, y),
		hitbox:   internal.NewBoundingBox(x, y, width, height),
		Velocity: velocity,
		bounds:   boundedTo,
	}
}

func (b *Bullet) Move(direction int) {
	b.pos.Y += int32(float32(direction) * b.Velocity)
	b.hitbox = internal.NewBoundingBox(b.pos.X, b.pos.Y, width, height)
}

func (b *Bullet) Bounds() internal.BoundingBox {
	return b.hitbox
}

func (b *Bullet) Position() internal.Point {
	return b.pos
}

func (b *Bullet) IsOutOfBounds() bool {
	return !b.bounds.Contains(b.pos)
}

func (b *Bullet) Collided() {
	b.collided = true
}

func (b *Bullet) Draw() {
	rl.DrawRectangle(b.pos.X, b.pos.Y, width, height, rl.Yellow)
}

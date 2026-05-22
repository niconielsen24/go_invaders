package mobs

import (
	"go_invaders/game/bullet"
	"go_invaders/game/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const greenStepInterval = 1.2

type Green struct {
	Sprite      rl.Texture2D
	BoundingBox internal.BoundingBox
	Position    internal.Point
	gameBounds  internal.BoundingBox
	stepTimer   float32
	directionX  int32
	alive       bool
	hitPoints   int
}

func NewGreen(sprite rl.Texture2D, position internal.Point, gameBounds internal.BoundingBox) *Green {
	return &Green{
		Sprite:      sprite,
		BoundingBox: centeredBox(position, spriteSize),
		Position:    position,
		gameBounds:  gameBounds,
		directionX:  1,
		alive:       true,
		hitPoints:   1,
	}
}

// Sweeps horizontally, drops a row and reverses when hitting the edge.
func (g *Green) Move() {
	g.stepTimer += rl.GetFrameTime()
	if g.stepTimer < greenStepInterval {
		return
	}
	g.stepTimer = 0

	next := g.Position.X + 15*g.directionX
	if next < g.gameBounds.X || next+spriteSize > g.gameBounds.X+g.gameBounds.Width {
		g.Position.Y += spriteSize
		g.directionX *= -1
	} else {
		g.Position.X = next
	}
	g.BoundingBox = centeredBox(g.Position, spriteSize)
}

func (g *Green) Shoot() *bullet.Bullet { return nil }

func (g *Green) CollidesWith(other internal.BoundingBox) bool {
	return g.BoundingBox.Intersects(other)
}

func (g *Green) Draw() {
	rl.DrawTexture(g.Sprite, g.Position.X, g.Position.Y, rl.Green)
}

func (g *Green) IsAlive() bool { return g.alive }

func (g *Green) Hit() {
	g.hitPoints--
	if g.hitPoints <= 0 {
		g.alive = false
	}
}

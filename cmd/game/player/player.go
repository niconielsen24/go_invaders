package player

import (
	"go_invaders/game/bullet"
	"go_invaders/game/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	sprite   rl.Texture2D
	position internal.Point
	bounds   internal.BoundingBox
}

func NewPlayer(bondedTo internal.BoundingBox, startingPosition internal.Point) *Player {
	return &Player{
		sprite:   rl.LoadTexture("assets/player.png"),
		position: startingPosition,
		bounds:   bondedTo,
	}
}

// INITIALIZE and POSITIONING logic.
func (p *Player) Destroy() {
	rl.UnloadTexture(p.sprite)
}

func (p *Player) Draw() {
	rl.DrawTexture(p.sprite, p.position.X, p.position.Y, rl.White)
}

func (p *Player) UpdatePosition(dx, dy int32) {
	newX := p.position.X + dx
	newY := p.position.Y + dy

	if newX < p.bounds.X {
		newX = p.bounds.X
	} else if newX > p.bounds.X+p.bounds.Width {
		newX = p.bounds.X + p.bounds.Width
	}

	if newY < p.bounds.Y {
		newY = p.bounds.Y
	} else if newY > p.bounds.Y+p.bounds.Height {
		newY = p.bounds.Y + p.bounds.Height
	}

	p.position = internal.NewPoint(newX, newY)
}

func (p *Player) GetPosition() internal.Point {
	return p.position
}

// SHOOTING logic.
func (p *Player) Shoot() bullet.Bullet {
	return bullet.NewBullet(p.position.X+25, p.position.Y, 10, p.bounds)
}

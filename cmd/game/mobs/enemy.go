package mobs

import (
	"go_invaders/game/bullet"
	"go_invaders/game/internal"
)

type Move interface {
	Move()
}

type Shoot interface {
	Shoot() *bullet.Bullet
}

type Hit interface {
	Hit()
}

type Life interface {
	IsAlive() bool
}

type Collider interface {
	CollidesWith(other internal.BoundingBox) bool
}

type Enemy interface {
	Move
	Shoot
	Hit
	Life
	Collider
}

package game

import (
	"go_invaders/game/bullet"
	"go_invaders/game/errors"
	"go_invaders/game/input"
	"go_invaders/game/internal"
	"go_invaders/game/mobs"
	"go_invaders/game/player"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type sprites = map[string]rl.Texture2D

type enemies = map[int]mobs.DrawableEnemy
type bullets = []bullet.Bullet

const (
	upperBound = 20
	leftBound  = 20
)

type Game struct {
	windowWidth  int32
	windowHeight int32
	initialized  bool
	destroyed    bool
	enemies      enemies
	deadEnemies  enemies
	player       *player.Player
	inputParser  *input.InputParser
	bounds       internal.BoundingBox
	liveBullets  bullets
	deadBullets  bullets
	sprites      sprites
}

func New(windowWidth, windowHeight int32) *Game {
	return &Game{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		initialized:  false,
		destroyed:    false,
		enemies:      nil,
		deadEnemies:  nil,
		inputParser:  nil,
		bounds: internal.NewBoundingBox(
			upperBound, leftBound,
			windowWidth-40, windowHeight-40,
		),
	}
}

func (g *Game) Run() error {
	if !g.initialized {
		return errors.ErrGameNotInitialized
	}

	if g.destroyed {
		return errors.ErrGameAlreadyDestroyed
	}

	for !rl.WindowShouldClose() {
		// Keep player intent first.
		input := g.inputParser.Parse()
		g.handleInput(input)

		//Compute collisions against previous frame state.
		g.computeCollisions()

		// Update game state.
		g.updateEnemies()
		g.updateBullets()
		g.destroyBullets()

		// Render.
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		g.drawEnemies()
		g.player.Draw()
		g.drawBullets()

		rl.EndDrawing()
	}

	return nil
}

func (g *Game) Init() {
	rl.InitWindow(g.windowWidth, g.windowHeight, "Go Invaders")
	rl.SetTargetFPS(60)
	g.sprites = sprites{
		"red":    rl.LoadTexture("assets/img/red.png"),
		"green":  rl.LoadTexture("assets/img/green.png"),
		"yellow": rl.LoadTexture("assets/img/yellow.png"),
		"extra":  rl.LoadTexture("assets/img/extra.png"),
	}
	g.initEnemies()
	g.deadEnemies = make(enemies)
	g.player = player.NewPlayer(
		g.bounds,
		internal.NewPoint(g.windowWidth/2, g.windowHeight-50),
	)
	g.inputParser = input.NewInputParser()
	g.initialized = true
}

func (g *Game) Destroy() {
	for _, sprite := range g.sprites {
		rl.UnloadTexture(sprite)
	}
	g.player.Destroy()
	rl.CloseWindow()
	g.destroyed = true
}

func (g *Game) initEnemies() {
	const cols = 16
	spacing := g.bounds.Width / cols
	row := func(y, col int) internal.Point {
		return internal.NewPoint(int32(col)*spacing+spacing/2, int32(y))
	}

	g.enemies = make(enemies)
	id := 0
	for i := range cols {
		g.enemies[id] = mobs.NewGreen(g.sprites["green"], row(150, i), g.bounds)
		id++
		g.enemies[id] = mobs.NewRed(g.sprites["red"], row(250, i))
		id++
		g.enemies[id] = mobs.NewYellow(g.sprites["yellow"], row(350, i))
		id++
	}
	g.enemies[id] = mobs.NewExtra(g.sprites["extra"], internal.NewPoint(0, 50), g.bounds)
}

func (g *Game) computeCollisions() {
	g.liveBullets = slices.DeleteFunc(g.liveBullets, func(b bullet.Bullet) bool {
		for id, enemy := range g.enemies {
			if enemy.CollidesWith(b.Bounds()) {
				enemy.Hit()
				if !enemy.IsAlive() {
					g.deadEnemies[id] = enemy
					delete(g.enemies, id)
				}
				return true
			}
		}
		return false
	})
}

func (g *Game) updateEnemies() {
	for _, enemy := range g.enemies {
		enemy.Move()
	}
}

func (g *Game) drawEnemies() {
	for _, enemy := range g.enemies {
		enemy.Draw()
	}
}

func (g *Game) handleInput(in int) {
	switch in {
	case input.MOVE_LEFT:
		g.player.UpdatePosition(-5, 0)
	case input.MOVE_RIGHT:
		g.player.UpdatePosition(5, 0)
	case input.SPACE_BAR:
		g.liveBullets = append(g.liveBullets, g.player.Shoot())
	}
}

func (g *Game) destroyBullets() {
	g.liveBullets = slices.DeleteFunc(g.liveBullets, func(p bullet.Bullet) bool {
		return p.IsOutOfBounds()
	})
}

func (g *Game) updateBullets() {
	for i := range g.liveBullets {
		g.liveBullets[i].Move(bullet.Up)
	}
}

func (g *Game) drawBullets() {
	for _, bull := range g.liveBullets {
		bull.Draw()
	}
}

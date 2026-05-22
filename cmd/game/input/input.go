package input

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	MOVE_LEFT = iota + 1
	MOVE_RIGHT
	SPACE_BAR
)

type InputParser struct{}

func NewInputParser() *InputParser {
	return &InputParser{}
}

func (p *InputParser) Parse() int {

	if rl.IsKeyDown(rl.KeyA) {
		return MOVE_LEFT
	}
	if rl.IsKeyDown(rl.KeyD) {
		return MOVE_RIGHT
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		return SPACE_BAR
	}

	return 0
}

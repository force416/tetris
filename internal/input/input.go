package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Action int

const (
	ActionMoveLeft Action = iota
	ActionMoveRight
	ActionRotate
	ActionSoftDrop
	ActionTogglePause
	ActionRestart
)

type Poller struct{}

func NewPoller() *Poller {
	return &Poller{}
}

func (p *Poller) Actions() []Action {
	actions := make([]Action, 0, 6)

	if justPressedOrRepeating(ebiten.KeyArrowLeft, 10, 4) || justPressedOrRepeating(ebiten.KeyA, 10, 4) {
		actions = append(actions, ActionMoveLeft)
	}
	if justPressedOrRepeating(ebiten.KeyArrowRight, 10, 4) || justPressedOrRepeating(ebiten.KeyD, 10, 4) {
		actions = append(actions, ActionMoveRight)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		actions = append(actions, ActionRotate)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		actions = append(actions, ActionSoftDrop)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		actions = append(actions, ActionTogglePause)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		actions = append(actions, ActionRestart)
	}

	return actions
}

func justPressedOrRepeating(key ebiten.Key, delay, interval int) bool {
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

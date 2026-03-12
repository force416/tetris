package app

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"tetris/internal/game"
	"tetris/internal/input"
	"tetris/internal/render"
)

type App struct {
	game        *game.Game
	input       *input.Poller
	lastUpdate  time.Time
	accumulator time.Duration
}

func New(g *game.Game) *App {
	return &App{
		game:       g,
		input:      input.NewPoller(),
		lastUpdate: time.Now(),
	}
}

func (a *App) Run() error {
	width, height := render.WindowSize(a.game.Config)
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Tetris")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	return ebiten.RunGame(a)
}

func (a *App) Update() error {
	now := time.Now()
	delta := now.Sub(a.lastUpdate)
	a.lastUpdate = now
	if delta > 250*time.Millisecond {
		delta = 250 * time.Millisecond
	}
	a.accumulator += delta

	for _, action := range a.input.Actions() {
		switch action {
		case input.ActionMoveLeft:
			a.game.HandleCommand(game.CommandMoveLeft)
		case input.ActionMoveRight:
			a.game.HandleCommand(game.CommandMoveRight)
		case input.ActionRotate:
			a.game.HandleCommand(game.CommandRotateCW)
		case input.ActionSoftDrop:
			a.game.HandleCommand(game.CommandSoftDrop)
		case input.ActionRestart:
			a.game.HandleCommand(game.CommandRestart)
			a.accumulator = 0
		case input.ActionTogglePause:
			if a.game.Status == game.StatusPaused {
				a.game.HandleCommand(game.CommandResume)
			} else if a.game.Status == game.StatusRunning {
				a.game.HandleCommand(game.CommandPause)
			}
		}
	}

	for a.accumulator >= a.game.GravityInterval() {
		a.game.Tick()
		a.accumulator -= a.game.GravityInterval()
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	render.Draw(screen, a.game)
}

func (a *App) Layout(_, _ int) (int, int) {
	return render.WindowSize(a.game.Config)
}

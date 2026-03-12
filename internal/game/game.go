package game

import "time"

type Status string

const (
	StatusRunning  Status = "running"
	StatusPaused   Status = "paused"
	StatusGameOver Status = "game_over"
)

type Command int

const (
	CommandMoveLeft Command = iota
	CommandMoveRight
	CommandRotateCW
	CommandSoftDrop
	CommandPause
	CommandResume
	CommandRestart
)

type Game struct {
	Config       Config
	Board        [][]TetrominoKind
	Active       Piece
	NextKind     TetrominoKind
	Score        int
	Level        int
	LinesCleared int
	Status       Status

	sequence Sequence
}

func New(config Config, sequence Sequence) *Game {
	if sequence == nil {
		sequence = NewBagSequence(0)
	}

	g := &Game{
		Config:   config,
		sequence: sequence,
	}
	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.Board = make([][]TetrominoKind, g.Config.BoardHeight)
	for y := range g.Board {
		g.Board[y] = make([]TetrominoKind, g.Config.BoardWidth)
	}

	g.Score = 0
	g.Level = g.Config.StartLevel
	g.LinesCleared = 0
	g.Status = StatusRunning
	g.NextKind = g.sequence.Next()
	g.spawnNextPiece()
}

func (g *Game) HandleCommand(command Command) {
	switch command {
	case CommandRestart:
		g.Restart()
	case CommandPause:
		if g.Status == StatusRunning {
			g.Status = StatusPaused
		}
	case CommandResume:
		if g.Status == StatusPaused {
			g.Status = StatusRunning
		}
	case CommandMoveLeft:
		if g.Status == StatusRunning {
			g.moveActive(-1, 0, g.Active.Rotation)
		}
	case CommandMoveRight:
		if g.Status == StatusRunning {
			g.moveActive(1, 0, g.Active.Rotation)
		}
	case CommandRotateCW:
		if g.Status == StatusRunning {
			g.rotateActive()
		}
	case CommandSoftDrop:
		if g.Status == StatusRunning && !g.moveActive(0, 1, g.Active.Rotation) {
			g.lockActivePiece()
		}
	}
}

func (g *Game) Tick() {
	if g.Status != StatusRunning {
		return
	}
	if g.moveActive(0, 1, g.Active.Rotation) {
		return
	}
	g.lockActivePiece()
}

func (g *Game) GravityInterval() time.Duration {
	return g.Config.GravityForLevel(g.Level)
}

func (g *Game) CellKind(x, y int) TetrominoKind {
	if y < 0 || y >= len(g.Board) || x < 0 || x >= len(g.Board[y]) {
		return EmptyKind
	}

	if g.Status != StatusGameOver {
		for _, cell := range g.Active.Cells() {
			if cell.X == x && cell.Y == y {
				return g.Active.Kind
			}
		}
	}

	return g.Board[y][x]
}

func (g *Game) moveActive(dx, dy, rotation int) bool {
	candidate := Piece{
		Kind:     g.Active.Kind,
		Rotation: rotation,
		X:        g.Active.X + dx,
		Y:        g.Active.Y + dy,
	}

	if !g.canPlace(candidate) {
		return false
	}

	g.Active = candidate
	return true
}

func (g *Game) rotateActive() bool {
	nextRotation := (g.Active.Rotation + 1) % len(pieceRotations[g.Active.Kind])
	for _, offset := range g.Config.WallKickOffsets {
		candidate := Piece{
			Kind:     g.Active.Kind,
			Rotation: nextRotation,
			X:        g.Active.X + offset.X,
			Y:        g.Active.Y + offset.Y,
		}
		if g.canPlace(candidate) {
			g.Active = candidate
			return true
		}
	}
	return false
}

func (g *Game) canPlace(piece Piece) bool {
	for _, cell := range piece.Cells() {
		if cell.X < 0 || cell.X >= g.Config.BoardWidth || cell.Y < 0 || cell.Y >= g.Config.BoardHeight {
			return false
		}
		if g.Board[cell.Y][cell.X] != EmptyKind {
			return false
		}
	}
	return true
}

func (g *Game) lockActivePiece() {
	for _, cell := range g.Active.Cells() {
		g.Board[cell.Y][cell.X] = g.Active.Kind
	}

	lines := g.clearLines()
	g.updateScoreAndLevel(lines)
	g.spawnNextPiece()
}

func (g *Game) clearLines() int {
	cleared := 0
	compacted := make([][]TetrominoKind, 0, len(g.Board))

	for _, row := range g.Board {
		full := true
		for _, kind := range row {
			if kind == EmptyKind {
				full = false
				break
			}
		}

		if full {
			cleared++
			continue
		}

		copyRow := append([]TetrominoKind(nil), row...)
		compacted = append(compacted, copyRow)
	}

	for len(compacted) < g.Config.BoardHeight {
		compacted = append([][]TetrominoKind{make([]TetrominoKind, g.Config.BoardWidth)}, compacted...)
	}

	g.Board = compacted
	return cleared
}

func (g *Game) updateScoreAndLevel(lines int) {
	if lines <= 0 {
		return
	}

	g.LinesCleared += lines
	g.Score += g.Config.ScoreByLines[lines] * g.Level
	g.Level = g.Config.StartLevel + (g.LinesCleared / g.Config.LinesPerLevel)
}

func (g *Game) spawnNextPiece() {
	g.Active = Piece{
		Kind:     g.NextKind,
		Rotation: 0,
		X:        g.Config.SpawnX,
		Y:        g.Config.SpawnY,
	}
	g.NextKind = g.sequence.Next()

	if !g.canPlace(g.Active) {
		g.Status = StatusGameOver
	}
}

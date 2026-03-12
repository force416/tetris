package game

import "testing"

func newTestGame() *Game {
	config := DefaultConfig()
	return New(config, NewFixedSequence([]TetrominoKind{TKind, IKind, OKind, LKind}))
}

func TestRotatePieceInOpenSpace(t *testing.T) {
	g := newTestGame()
	g.Active = Piece{Kind: TKind, Rotation: 0, X: 3, Y: 0}

	g.HandleCommand(CommandRotateCW)

	if g.Active.Rotation != 1 {
		t.Fatalf("expected rotation to advance to 1, got %d", g.Active.Rotation)
	}

	expected := []Point{{X: 4, Y: 0}, {X: 4, Y: 1}, {X: 5, Y: 1}, {X: 4, Y: 2}}
	actual := g.Active.Cells()
	for index := range expected {
		if actual[index] != expected[index] {
			t.Fatalf("expected rotated cell %d to be %+v, got %+v", index, expected[index], actual[index])
		}
	}
}

func TestRejectMoveIntoSettledBlocks(t *testing.T) {
	g := newTestGame()
	g.Active = Piece{Kind: OKind, Rotation: 0, X: 3, Y: 0}
	g.Board[0][3] = JKind
	g.Board[1][3] = JKind

	original := g.Active
	g.HandleCommand(CommandMoveLeft)

	if g.Active != original {
		t.Fatalf("expected piece to stay in place when moving into settled blocks, got %+v", g.Active)
	}
}

func TestClearLineAndIncreaseScore(t *testing.T) {
	g := New(DefaultConfig(), NewFixedSequence([]TetrominoKind{IKind, OKind, TKind}))
	g.Active = Piece{Kind: IKind, Rotation: 0, X: 0, Y: 17}
	g.NextKind = OKind

	for x := 4; x < g.Config.BoardWidth; x++ {
		g.Board[19][x] = LKind
	}

	g.HandleCommand(CommandSoftDrop)
	g.HandleCommand(CommandSoftDrop)

	if g.LinesCleared != 1 {
		t.Fatalf("expected 1 cleared line, got %d", g.LinesCleared)
	}
	if g.Score != 100 {
		t.Fatalf("expected score to be 100 after single line clear, got %d", g.Score)
	}
	if g.Board[19][0] != EmptyKind {
		t.Fatalf("expected cleared row to be empty after line removal")
	}
	if g.Active.Kind != OKind {
		t.Fatalf("expected next piece to become active after lock, got %s", g.Active.Kind)
	}
}

func TestSpawnBlockedTriggersGameOver(t *testing.T) {
	g := newTestGame()
	g.NextKind = OKind
	g.Board[0][4] = ZKind

	g.spawnNextPiece()

	if g.Status != StatusGameOver {
		t.Fatalf("expected game over when spawn area is blocked, got %s", g.Status)
	}
}

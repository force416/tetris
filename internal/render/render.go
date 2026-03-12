package render

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"tetris/internal/game"
)

const (
	CellSize      = 28
	BoardPadding  = 24
	SidebarWidth  = 220
	PreviewOffset = 120
)

var (
	backgroundColor = color.RGBA{R: 18, G: 24, B: 38, A: 255}
	panelColor      = color.RGBA{R: 28, G: 36, B: 54, A: 255}
	gridColor       = color.RGBA{R: 45, G: 58, B: 84, A: 255}
	overlayColor    = color.RGBA{R: 9, G: 12, B: 18, A: 210}
)

func WindowSize(config game.Config) (int, int) {
	width := BoardPadding*2 + (config.BoardWidth * CellSize) + SidebarWidth
	height := BoardPadding*2 + (config.BoardHeight * CellSize)
	return width, height
}

func Draw(screen *ebiten.Image, g *game.Game) {
	screen.Fill(backgroundColor)
	drawBoard(screen, g)
	drawSidebar(screen, g)
	drawStatus(screen, g)
}

func drawBoard(screen *ebiten.Image, g *game.Game) {
	boardWidth := float32(g.Config.BoardWidth * CellSize)
	boardHeight := float32(g.Config.BoardHeight * CellSize)
	vector.DrawFilledRect(screen, BoardPadding, BoardPadding, boardWidth, boardHeight, panelColor, false)

	for y := 0; y < g.Config.BoardHeight; y++ {
		for x := 0; x < g.Config.BoardWidth; x++ {
			px := BoardPadding + (x * CellSize)
			py := BoardPadding + (y * CellSize)
			vector.StrokeRect(screen, float32(px), float32(py), CellSize, CellSize, 1, gridColor, false)
			drawBlock(screen, px+2, py+2, CellSize-4, g.CellKind(x, y))
		}
	}
}

func drawSidebar(screen *ebiten.Image, g *game.Game) {
	x := BoardPadding*2 + (g.Config.BoardWidth * CellSize)
	width := SidebarWidth - BoardPadding
	height := BoardPadding*2 + (g.Config.BoardHeight * CellSize)
	vector.DrawFilledRect(screen, float32(x), 0, float32(width), float32(height), panelColor, false)

	ebitenutil.DebugPrintAt(screen, "TETRIS", x+20, BoardPadding)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("分數: %d", g.Score), x+20, BoardPadding+40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("等級: %d", g.Level), x+20, BoardPadding+60)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("消行: %d", g.LinesCleared), x+20, BoardPadding+80)
	ebitenutil.DebugPrintAt(screen, "下一個方塊", x+20, BoardPadding+120)

	drawPreview(screen, x+40, BoardPadding+PreviewOffset, g.NextKind)

	ebitenutil.DebugPrintAt(screen, "控制:", x+20, BoardPadding+260)
	ebitenutil.DebugPrintAt(screen, "← → 移動", x+20, BoardPadding+280)
	ebitenutil.DebugPrintAt(screen, "↑ / Space 旋轉", x+20, BoardPadding+300)
	ebitenutil.DebugPrintAt(screen, "↓ 加速下落", x+20, BoardPadding+320)
	ebitenutil.DebugPrintAt(screen, "P 暫停", x+20, BoardPadding+340)
	ebitenutil.DebugPrintAt(screen, "R 重開", x+20, BoardPadding+360)
}

func drawPreview(screen *ebiten.Image, left, top int, kind game.TetrominoKind) {
	for _, point := range game.PieceCells(kind, 0) {
		drawBlock(screen, left+(point.X*20), top+(point.Y*20), 18, kind)
	}
}

func drawStatus(screen *ebiten.Image, g *game.Game) {
	if g.Status == game.StatusRunning {
		return
	}

	width, height := screen.Bounds().Dx(), screen.Bounds().Dy()
	vector.DrawFilledRect(screen, 0, 0, float32(width), float32(height), overlayColor, false)

	message := "已暫停"
	subtitle := "按 P 繼續"
	if g.Status == game.StatusGameOver {
		message = "遊戲結束"
		subtitle = "按 R 重新開始"
	}

	ebitenutil.DebugPrintAt(screen, message, width/2-32, height/2-10)
	ebitenutil.DebugPrintAt(screen, subtitle, width/2-46, height/2+14)
}

func drawBlock(screen *ebiten.Image, x, y, size int, kind game.TetrominoKind) {
	if kind == game.EmptyKind {
		return
	}
	base := colorForKind(kind)
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(size), float32(size), base, false)
	highlight := lighten(base, 24)
	shadow := darken(base, 24)
	vector.StrokeRect(screen, float32(x), float32(y), float32(size), float32(size), 2, highlight, false)
	vector.StrokeLine(screen, float32(x), float32(y+size-1), float32(x+size-1), float32(y+size-1), 2, shadow, false)
	vector.StrokeLine(screen, float32(x+size-1), float32(y), float32(x+size-1), float32(y+size-1), 2, shadow, false)
}

func colorForKind(kind game.TetrominoKind) color.RGBA {
	switch kind {
	case game.IKind:
		return color.RGBA{R: 34, G: 211, B: 238, A: 255}
	case game.OKind:
		return color.RGBA{R: 250, G: 204, B: 21, A: 255}
	case game.TKind:
		return color.RGBA{R: 192, G: 132, B: 252, A: 255}
	case game.SKind:
		return color.RGBA{R: 74, G: 222, B: 128, A: 255}
	case game.ZKind:
		return color.RGBA{R: 248, G: 113, B: 113, A: 255}
	case game.JKind:
		return color.RGBA{R: 96, G: 165, B: 250, A: 255}
	case game.LKind:
		return color.RGBA{R: 251, G: 146, B: 60, A: 255}
	default:
		return color.RGBA{}
	}
}

func lighten(base color.RGBA, offset uint8) color.RGBA {
	return color.RGBA{
		R: saturatingAdd(base.R, offset),
		G: saturatingAdd(base.G, offset),
		B: saturatingAdd(base.B, offset),
		A: base.A,
	}
}

func darken(base color.RGBA, offset uint8) color.RGBA {
	return color.RGBA{
		R: saturatingSub(base.R, offset),
		G: saturatingSub(base.G, offset),
		B: saturatingSub(base.B, offset),
		A: base.A,
	}
}

func saturatingAdd(value, delta uint8) uint8 {
	sum := int(value) + int(delta)
	if sum > 255 {
		return 255
	}
	return uint8(sum)
}

func saturatingSub(value, delta uint8) uint8 {
	diff := int(value) - int(delta)
	if diff < 0 {
		return 0
	}
	return uint8(diff)
}

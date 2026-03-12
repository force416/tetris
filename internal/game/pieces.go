package game

type TetrominoKind string

const (
	EmptyKind TetrominoKind = ""
	IKind     TetrominoKind = "I"
	OKind     TetrominoKind = "O"
	TKind     TetrominoKind = "T"
	SKind     TetrominoKind = "S"
	ZKind     TetrominoKind = "Z"
	JKind     TetrominoKind = "J"
	LKind     TetrominoKind = "L"
)

type Point struct {
	X int
	Y int
}

type Piece struct {
	Kind     TetrominoKind
	Rotation int
	X        int
	Y        int
}

var allKinds = []TetrominoKind{IKind, OKind, TKind, SKind, ZKind, JKind, LKind}

var pieceRotations = map[TetrominoKind][][]Point{
	IKind: {
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}},
		{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}},
		{{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 2}, {X: 3, Y: 2}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 1, Y: 3}},
	},
	OKind: {
		{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}},
	},
	TKind: {
		{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
		{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}},
	},
	SKind: {
		{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 2}},
		{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}},
	},
	ZKind: {
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 2, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 2}},
		{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 0, Y: 2}},
	},
	JKind: {
		{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}},
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 2}},
	},
	LKind: {
		{{X: 2, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 2}},
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 0, Y: 2}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}},
	},
}

func PieceCells(kind TetrominoKind, rotation int) []Point {
	rotations := pieceRotations[kind]
	if len(rotations) == 0 {
		return nil
	}
	index := rotation % len(rotations)
	if index < 0 {
		index += len(rotations)
	}
	return rotations[index]
}

func (p Piece) Cells() []Point {
	local := PieceCells(p.Kind, p.Rotation)
	cells := make([]Point, 0, len(local))
	for _, point := range local {
		cells = append(cells, Point{X: p.X + point.X, Y: p.Y + point.Y})
	}
	return cells
}

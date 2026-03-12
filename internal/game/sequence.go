package game

import (
	"math/rand"
	"time"
)

type Sequence interface {
	Next() TetrominoKind
}

type BagSequence struct {
	random *rand.Rand
	bag    []TetrominoKind
}

func NewBagSequence(seed int64) *BagSequence {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	return &BagSequence{
		random: rand.New(rand.NewSource(seed)),
	}
}

func (b *BagSequence) refill() {
	b.bag = append([]TetrominoKind(nil), allKinds...)
	b.random.Shuffle(len(b.bag), func(i, j int) {
		b.bag[i], b.bag[j] = b.bag[j], b.bag[i]
	})
}

func (b *BagSequence) Next() TetrominoKind {
	if len(b.bag) == 0 {
		b.refill()
	}
	next := b.bag[0]
	b.bag = b.bag[1:]
	return next
}

type FixedSequence struct {
	kinds    []TetrominoKind
	index    int
	fallback TetrominoKind
}

// NewFixedSequence keeps piece generation deterministic for tests and scripted scenarios.
func NewFixedSequence(kinds []TetrominoKind) *FixedSequence {
	return &FixedSequence{
		kinds:    append([]TetrominoKind(nil), kinds...),
		fallback: IKind,
	}
}

func (f *FixedSequence) Next() TetrominoKind {
	if len(f.kinds) == 0 {
		return f.fallback
	}
	if f.index >= len(f.kinds) {
		return f.kinds[len(f.kinds)-1]
	}
	next := f.kinds[f.index]
	f.index++
	return next
}

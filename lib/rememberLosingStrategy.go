package tictacgo

import (
	//"fmt"
	"math/rand"
)

//RememberLosingStrategy avoids things it picked when it lost
type RememberLosingStrategy struct {
	enabled  bool
	name     string
	memories map[uint][]byte
	moves    map[uint][]byte
}

//NewRememberLosingStrategy constructor
func NewRememberLosingStrategy() *RememberLosingStrategy {
	return &RememberLosingStrategy{
		name:     "RememberLosingStrategy",
		enabled:  true,
		memories: make(map[uint][]byte),
		moves:    make(map[uint][]byte),
	}
}

func (r *RememberLosingStrategy) resetMoves() {
	r.moves = make(map[uint][]byte)
}

// Name of the strategy
func (r *RememberLosingStrategy) Name() string {
	return r.name
}

// Move avoids things it picked when it lost
func (r *RememberLosingStrategy) Move(b Board) byte {
	var index int
	var choice byte
	emptySquares := b.ListEmptySquares()

	// compare to the memories
	boardNumber := uint(b.Int())
	availableSquares := filterBytes(emptySquares, r.memories[boardNumber])

	if len(availableSquares) != 0 {
		index = rand.Intn(len(availableSquares))
		choice = byte(availableSquares[index])
	} else {
		r.memories[boardNumber] = nil
		if len(emptySquares) != 0 {
			index = rand.Intn(len(emptySquares))
			choice = byte(emptySquares[index])
		} else {
			choice = 0
		}
	}

	// Now we jot down our move so we can think about it later
	r.moves[boardNumber] = append(r.moves[boardNumber], choice)
	return choice
}

// Win adjust strategy on a win
func (r *RememberLosingStrategy) Win(b Board) {
	r.resetMoves()
	return
}

// Lose adjust strategy on a lose
func (r *RememberLosingStrategy) Lose(b Board) {
	for board := range r.moves {
		for _, m := range r.moves[board] {
			r.memories[board] = append(r.memories[board], m)
		}
		r.memories[board] = uniqueBytes(r.memories[board])
	}
	r.resetMoves()
	return
}

// Draw adjust strategy on a draw
func (r *RememberLosingStrategy) Draw(b Board) {
	r.resetMoves()
	return
}

package tictacgo

import (
	"fmt"
	"math/rand"
)

//RememberWinningStrategy avoids things it picked when it lost
type RememberWinningStrategy struct {
	enabled  bool
	name     string
	memories map[uint][]byte
	moves    map[uint][]byte
}

//NewRememberWinningStrategy constructor
func NewRememberWinningStrategy() *RememberWinningStrategy {
	return &RememberWinningStrategy{
		name:     "RememberWinningStrategy",
		enabled:  true,
		memories: make(map[uint][]byte),
		moves:    make(map[uint][]byte),
	}
}

func (r *RememberWinningStrategy) resetMoves() {
	r.moves = make(map[uint][]byte)
}

// Name of the strategy
func (r *RememberWinningStrategy) Name() string {
	return r.name
}

// Move avoids things it picked when it lost
func (r *RememberWinningStrategy) Move(b Board) byte {
	var index int
	var choice byte
	boardNumber := uint(b.Int())
	emptySquares := b.ListEmptySquares()

	// add the squares you won with to increase their likelihood
	possibleSquares := append(emptySquares, r.memories[boardNumber]...)

	if len(possibleSquares) > 0 {
		index = rand.Intn(len(possibleSquares))
		choice = byte(possibleSquares[index])
	} else {
		choice = 0

	}

	// Now we jot down our move so we can think about it later
	r.moves[boardNumber] = append(r.moves[boardNumber], choice)
	return choice
}

// Win adjust strategy on a win
func (r *RememberWinningStrategy) Win(b Board) {
	fmt.Println(r.Name())
	for board := range r.moves {
		for _, m := range r.moves[board] {
			r.memories[board] = append(r.memories[board], m)
		}
		// by not unique this array, we remember more stronly
		maxLength := 10
		memLength := len(r.memories[board])
		if memLength > maxLength {
			sliceStart := memLength - maxLength
			r.memories[board] = r.memories[board][sliceStart:]
		}
	}
	r.resetMoves()
	return
}

// Lose adjust strategy on a lose
func (r *RememberWinningStrategy) Lose(b Board) {
	r.resetMoves()
	return
}

// Draw adjust strategy on a draw
func (r *RememberWinningStrategy) Draw(b Board) {
	r.resetMoves()
	return
}

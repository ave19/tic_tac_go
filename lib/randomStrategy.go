package tictacgo

import (
	"fmt"
	"math/rand"
)

// RandomStrategy just picks a random spot on the board
type RandomStrategy struct {
	enabled bool
	name    string
}

// NewRandomStrategy constructor
func NewRandomStrategy() *RandomStrategy {
	return &RandomStrategy{name: "RandomStrategy", enabled: true}
}

// Name of the strategy
func (r *RandomStrategy) Name() string {
	return r.name
}

// Enabled true or false
func (r *RandomStrategy) Enabled() bool {
	return r.enabled
}

// Move selects an empty square from the board
func (r *RandomStrategy) Move(b Board) byte {
	emptySquares := b.ListEmptySquares()
	//fmt.Println("emptySquares = ", emptySquares)
	//fmt.Println("emptySquareLen = ", len(emptySquares))
	if len(emptySquares) == 0 {
		return 0
	}
	choice := rand.Intn(len(emptySquares))
	fmt.Println("choosing square ", emptySquares[choice])
	return byte(emptySquares[choice])
}

// Win adjust strategy on a win
func (r *RandomStrategy) Win(b Board) {
	return
}

// Lose adjust strategy on a lose
func (r *RandomStrategy) Lose(b Board) {
	return
}

// Draw adjust strategy on a draw
func (r *RandomStrategy) Draw(b Board) {
	return
}

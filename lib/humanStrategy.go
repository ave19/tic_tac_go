package tictacgo

import (
//"fmt"
)

// HumanStrategy just picks a random spot on the board
type HumanStrategy struct {
	enabled bool
	name    string
}

// NewHumanStrategy constructor
func NewHumanStrategy() *HumanStrategy {
	return &HumanStrategy{name: "HumanStrategy", enabled: true}
}

// Name of the strategy
func (r *HumanStrategy) Name() string {
	return r.name
}

// Enabled true or false
func (r *HumanStrategy) Enabled() bool {
	return r.enabled
}

// Move selects an empty square from the board
func (r *HumanStrategy) Move(b Board) (choice byte) {
	choice = 0
	return choice
}

// Win adjust strategy on a win
func (r *HumanStrategy) Win(b Board) {
	return
}

// Lose adjust strategy on a lose
func (r *HumanStrategy) Lose(b Board) {
	return
}

// Draw adjust strategy on a draw
func (r *HumanStrategy) Draw(b Board) {
	return
}

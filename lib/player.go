package tictacgo

import (
	"fmt"
)

// Strategy of various move methods
type Strategy interface {
	Name() string
	Move(Board) byte
	Win(Board)
	Lose(Board)
	Draw(Board)
}

// Player ready player one
type Player struct {
	name      string
	number    byte
	strategy  Strategy
	moves     []byte
	winCount  uint
	loseCount uint
	drawCount uint
}

// NewPlayer returns a new player
func NewPlayer() Player {
	var p Player
	p.name = "FooBear"
	p.number = 0
	p.strategy = NewRandomStrategy()
	p.winCount = 0
	p.loseCount = 0
	p.drawCount = 0
	return p
}

// SetName sets the player's name
func (p *Player) SetName(newName string) {
	p.name = newName
}

// Name returns the name
func (p Player) Name() string {
	return p.name
}

// SetNumber sets the player's number
func (p *Player) SetNumber(newNumber byte) {
	p.number = newNumber
}

// Number returns the name
func (p Player) Number() byte {
	return p.number
}

// SetStrategy for the player
func (p *Player) SetStrategy(s Strategy) {
	p.strategy = s
}

// Strategy of the player
func (p Player) Strategy() Strategy {
	return p.strategy
}

// Move returns a square based on the board
func (p Player) Move(b Board) byte {
	return p.strategy.Move(b)
}

// Stats about the player's record
func (p *Player) Stats() {
	fmt.Println(p)
	fmt.Println("Name:", p.Name())
	fmt.Println("  Number:", p.Number())
	fmt.Println("  Wins:", p.winCount)
	fmt.Println("  Loses:", p.loseCount)
}

// Win handler
func (p *Player) Win(b Board) {
	fmt.Println(" old win count:", p.winCount)
	p.winCount++
	p.strategy.Win(b)
	fmt.Println(" new win count:", p.winCount)
	return
}

// Lose handler
func (p *Player) Lose(b Board) {
	p.loseCount++
	p.strategy.Lose(b)
}

// Draw handler
func (p *Player) Draw(b Board) {
	p.drawCount++
	p.strategy.Draw(b)
}

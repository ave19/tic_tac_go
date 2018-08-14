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
	lossCount uint
	drawCount uint
	moveChan  chan byte
}

// NewPlayer returns a new player
func NewPlayer() Player {
	var p Player
	p.name = "FooBear"
	p.number = 0
	p.strategy = NewRandomStrategy()
	p.winCount = 0
	p.lossCount = 0
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
	totalGames := p.winCount + p.lossCount + p.drawCount
	fmt.Println("Name:", p.Name())
	fmt.Println("  Player:", p.Number())
	fmt.Println("  Strategy:", p.Strategy().Name())
	fmt.Println("  Total Games:", totalGames)
	fmt.Println("  Wins:", p.winCount)
	fmt.Println("  Loses:", p.lossCount)
	fmt.Println("  Draws:", p.drawCount)
	if totalGames > 0 {
		winPercentage := 100 * float64(p.winCount) / float64(totalGames)
		fmt.Printf("  Win Percentage: %.2f%%\n", winPercentage)
		lossPercentage := 100 * float64(p.lossCount) / float64(totalGames)
		fmt.Printf("  Loss Percentage: %.2f%%\n", lossPercentage)
	}
}

// Win handler
func (p *Player) Win(b Board) {
	Trace.Printf("%v: win", p.Name())
	p.winCount++
	p.strategy.Win(b)
	return
}

// Lose handler
func (p *Player) Lose(b Board) {
	Trace.Printf("%v: lose", p.Name())
	p.lossCount++
	p.strategy.Lose(b)
}

// Draw handler
func (p *Player) Draw(b Board) {
	Trace.Printf("%v: draw", p.Name())
	p.drawCount++
	p.strategy.Draw(b)
}

// Play handler
func (p *Player) Play(bChan chan Board, mChan chan byte) {
	fmt.Printf("%v: Ready to play!\n", p.Name())
	for board := range bChan {
		Trace.Printf("%v: received board: %v\n", p.Name(), board.Int())
		myMove := p.Move(board)
		Trace.Printf("%v: pics %v.\n", p.Name(), myMove)
		mChan <- myMove
	}
}

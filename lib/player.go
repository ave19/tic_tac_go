package tictacgo

import (
	"bytes"
	"fmt"
	"io"
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
func (p Player) Stats(output io.Writer) {
	var out bytes.Buffer
	totalGames := p.winCount + p.lossCount + p.drawCount
	out.WriteString(fmt.Sprintf("Name: %v\n", p.Name()))
	out.WriteString(fmt.Sprintf("  Player: %d\n", p.Number()))
	out.WriteString(fmt.Sprintf("  Strategy: %v\n", p.Strategy().Name()))
	out.WriteString(fmt.Sprintf("  Total Games: %d\n", totalGames))
	out.WriteString(fmt.Sprintf("  Wins: %d\n", p.winCount))
	out.WriteString(fmt.Sprintf("  Loses: %d\n", p.lossCount))
	out.WriteString(fmt.Sprintf("  Draws: %d\n", p.drawCount))
	if totalGames > 0 {
		winPercentage := 100 * float64(p.winCount) / float64(totalGames)
		out.WriteString(fmt.Sprintf("  Win Percentage: %.2f%%\n", winPercentage))
		lossPercentage := 100 * float64(p.lossCount) / float64(totalGames)
		out.WriteString(fmt.Sprintf("  Loss Percentage: %.2f%%\n", lossPercentage))
	}
	output.Write(out.Bytes())
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
	Trace.Println("bChan:", bChan)
	Trace.Println("mChan:", mChan)
	for board := range bChan {
		Trace.Printf("%v: received board: %v\n", p.Name(), board.Int())
		myMove := p.Move(board)
		Trace.Printf("%v: picks %v.\n", p.Name(), myMove)
		mChan <- myMove
	}
}

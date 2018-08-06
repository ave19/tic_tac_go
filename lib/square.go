package tictacgo

import (
	"fmt"
)

type state byte

const (
	empty state = iota
	x
	o
)

// Square is one part of the board
type Square struct {
	state state
}

func (s Square) String() string {
	switch s.state {
	case x:
		return "X"
	case o:
		return "O"
	case empty:
		return " "
	default:
		return "?"
	}
}

// Byte renders the byte value of the Square
func (s Square) Byte() byte {
	return byte(s.state)
}

// SetFromString changes the state to the indicated mark
func (s *Square) SetFromString(newValue string) {
	switch newValue {
	case " ":
		s.state = empty
	case "X":
		s.state = x
	case "O":
		s.state = o
	default:
		fmt.Println("I don't like that value, I'm ignoring you.")
	}
}

// SetFromByte changes the state to the indicated mark
func (s *Square) SetFromByte(newValue byte) {
	switch newValue {
	case 0:
		s.state = empty
	case 1:
		s.state = x
	case 2:
		s.state = o
	default:
		fmt.Printf("I don't like that value '%v', I'm ignoring you.\n", newValue)
	}
}

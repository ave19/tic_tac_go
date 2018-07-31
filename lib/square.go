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
	case X:
		return "X"
	case O:
		return "O"
	case empty:
		return "."
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
	fmt.Printf("&s = %p\n", &s)
	fmt.Printf("set from byte: new_value: %v\n", newValue)
	switch new_value {
	case 0:
		s.state = empty
	case 1:
		s.state = X
	case 2:
		s.state = O
	default:
		fmt.Printf("I don't like that value '%v', I'm ignoring you.\n", new_value)
	}
	fmt.Printf("s.state: %v\n", s.state)
	fmt.Printf("&s = %p\n", &s)
}

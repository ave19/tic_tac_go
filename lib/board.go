package tictacgo

import (
	"bytes"
	//"strings"
	"fmt"
	"strconv"
)

type board struct {
	board map[uint8]Square
}

// NewBoard returns a board structure
func NewBoard() (board *Board) {
	bPtr := &Board{}
	bPtr.board = make(map[uint8]Square, 9)
	for i := uint8(0); i < 10; i++ {
		b_ptr.board[i] = Square{0}
	}
	return b_ptr
}

func (b Board) String() string {
	var output bytes.Buffer
	output.WriteString("start...")
	for i := uint8(0); i < uint8(len(b.board)); i++ {
		output.WriteString(strconv.Itoa(int(i)))
		output.WriteString(" -> ")
		output.WriteString(b.board[i].String())
		output.WriteString("\n")
	}
	output.WriteString("...end\n")
	return output.String()
}

/*
func (board Board) Base3() string {
  var output bytes.Buffer
  output.WriteString(strconv.FormatInt(int64(board.Upper_left.Byte()), 10))
  output.WriteString(strconv.FormatInt(int64(board.Upper_center.Byte()), 10))
  output.WriteString(strconv.FormatInt(int64(board.Upper_right.Byte()), 10))
  output.WriteString(strconv.FormatInt(int64(board.Middle_left.Byte()), 10))
  output.WriteString(strconv.FormatInt(int64(board.Middle_center.Byte()), 10))
  output.WriteString(strconv.FormatInt(int64(board.Middle_right.Byte()), 10))
  output.WriteString(strconv.FormatInt(int64(board.Lower_left.Byte()), 10))
  output.WriteString(strconv.FormatInt(int64(board.Lower_center.Byte()), 10))
  output.WriteString(strconv.FormatInt(int64(board.Lower_right.Byte()), 10))
  return fmt.Sprintf("%09s", output.String())
}
*/

/*
func (board Board) Int() uint64 {
  // turn the squares into a string
  x, _ := strconv.ParseUint(board.Base3(), 3, 64)
  return x
}
*/

/*
func (board *Board) SetFromInt(new_value uint64) {
  if new_value > 19862 {
    fmt.Println("new value is too big")
  } else {
    board.SetFromBase3(strconv.FormatUint(new_value, 3))
  }
}
*/

/*
func (board *Board) SetFromBase3(new_value string) {
  str_representation := fmt.Sprintf("%09s", new_value)

  ul, _ := strconv.ParseInt(str_representation[0:1], 10, 8)
  uc, _ := strconv.ParseInt(str_representation[1:2], 10, 8)
  ur, _ := strconv.ParseInt(str_representation[2:3], 10, 8)
  ml, _ := strconv.ParseInt(str_representation[3:4], 10, 8)
  mc, _ := strconv.ParseInt(str_representation[4:5], 10, 8)
  mr, _ := strconv.ParseInt(str_representation[5:6], 10, 8)
  ll, _ := strconv.ParseInt(str_representation[6:7], 10, 8)
  lc, _ := strconv.ParseInt(str_representation[7:8], 10, 8)
  lr, _ := strconv.ParseInt(str_representation[8:], 10, 8)


  board.Upper_left.SetFromByte(byte(ul))
  board.Upper_center.SetFromByte(byte(uc))
  board.Upper_right.SetFromByte(byte(ur))
  board.Middle_left.SetFromByte(byte(ml))
  board.Middle_center.SetFromByte(byte(mc))
  board.Middle_right.SetFromByte(byte(mr))
  board.Lower_left.SetFromByte(byte(ll))
  board.Lower_center.SetFromByte(byte(lc))
  board.Lower_right.SetFromByte(byte(lr))
}
*/

/*
func (board Board) GetSquareByInt(square int) *Square {
  switch square {
  case 0:
    return board.Upper_left
  case 1:
    return board.Upper_center
  case 2:
    return board.Upper_right
  case 3:
    return board.Middle_left
  case 4:
    return board.Middle_center
  case 5:
    return board.Middle_right
  case 6:
    return board.Lower_left
  case 7:
    return board.Lower_center
  case 8:
    return board.Lower_right
  default:
    fmt.Println("I am very mad")
    return board.Middle_center
  }
}
*/

// Move adjusts the board to a moved state
func (b *Board) Move(square uint8, mark byte) {
	fmt.Printf("square: %v\n", square)
	fmt.Printf("move: square: %T\n", b.board[square])
	fmt.Printf("move: mark: %v\n", mark)
	sPtr := b.board[square]
	sPtr.SetFromByte(mark)
	//fmt.Printf("move board:\n%v\n", board.String())
}

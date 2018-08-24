package tictacgo

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

// There is a limit to the board values
var maxBoardValue = uint64(math.Pow(3, 9)) - 1

var winCombos = [8][3]byte{
	{0, 1, 2}, // across the top
	{3, 4, 5}, // across the middle
	{6, 7, 8}, // across the bottom
	{0, 4, 8}, // diagnol down
	{6, 4, 2}, // diagnol up
	{0, 3, 6}, // down left
	{1, 4, 7}, // down middle
	{2, 5, 8}, // down right
}

// Board is the tic tac toe board
type Board struct {
	board      []Square
	NumSquares uint
}

// NewBoard returns a board structure
func NewBoard() Board {
	numSquares := uint(9)
	var b Board
	b.board = make([]Square, numSquares)
	b.NumSquares = numSquares
	for i := byte(0); i < byte(numSquares); i++ {
		b.board[i].state = 0
	}
	Trace.Println("Created new board.")
	return b
}

// String returns the board in a visual tic tac toe paradigm
func (b Board) String() string {
	Trace.Println("Call to String()")
	var output bytes.Buffer
	output.WriteString(
		fmt.Sprintf(
			"%v|%v|%v\n-+-+-\n%v|%v|%v\n-+-+-\n%v|%v|%v\n",
			b.board[0].String(),
			b.board[1].String(),
			b.board[2].String(),
			b.board[3].String(),
			b.board[4].String(),
			b.board[5].String(),
			b.board[6].String(),
			b.board[7].String(),
			b.board[8].String()))
	return output.String()
}

// Base3 returns the base 3 string for the board state
func (b Board) Base3() string {
	var output bytes.Buffer
	for _, x := range b.board {
		output.WriteString(
			//fmt.Sprintf("%d", strconv.FormatInt(int64(x.Byte()), 10)))
			fmt.Sprintf("%d", x.Byte()))
	}
	/*
		strconv.FormatInt(int64(b.board[0].Byte()), 10),
		strconv.FormatInt(int64(b.board[1].Byte()), 10),
		strconv.FormatInt(int64(b.board[2].Byte()), 10),
		strconv.FormatInt(int64(b.board[3].Byte()), 10),
		strconv.FormatInt(int64(b.board[4].Byte()), 10),
		strconv.FormatInt(int64(b.board[5].Byte()), 10),
		strconv.FormatInt(int64(b.board[6].Byte()), 10),
		strconv.FormatInt(int64(b.board[7].Byte()), 10),
		strconv.FormatInt(int64(b.board[8].Byte()), 10))
	*/
	return fmt.Sprintf("%09s", output.String())
}

// Int returns the integer for the board state
func (b Board) Int() uint64 {
	// turn the squares into a string
	x, _ := strconv.ParseUint(b.Base3(), 3, 64)
	return x
}

// SetFromInt you give it an integer, it makes a board out of it
func (b *Board) SetFromInt(newValue uint64) (err error) {
	err = nil
	if newValue > maxBoardValue {
		err := ValueError{"New value is too big"}
		Error.Println("Cannot set:", err.message, newValue)
	} else {
		b.SetFromBase3(strconv.FormatUint(newValue, 3))
	}
	return err
}

// SetFromBase3 you give it a string like '010212120'
func (b *Board) SetFromBase3(newValue string) (err error) {
	err = nil
	strRepresentation := fmt.Sprintf("%09s", newValue)
	for i := 0; i < 9; i++ {
		x, err := strconv.ParseInt(strRepresentation[i:i+1], 10, 8)
		if err != nil {
			Error.Printf("%v: %v\n", err, x)
		}
		b.board[i].SetFromByte(byte(x))
	}
	return
}

// ListEmptySquares still available
func (b Board) ListEmptySquares() (emptySquares []byte) {
	for i := uint(0); i < b.NumSquares; i++ {
		if b.board[i].Byte() == 0 {
			emptySquares = append(emptySquares, byte(i))
		}
	}
	Trace.Printf("%d empty squares.\n", len(emptySquares))
	return
}

// Move adjusts the board to a moved state
func (b *Board) Move(square uint8, mark byte) (err error) {
	Trace.Printf("Move: square: %d, mark: %d\n", square, mark)
	err = b.board[square].SetFromByte(mark)
	return err
}

// Winner return the winning player if there is one
func (b *Board) Winner() (win bool, winner byte) {
	win = false
	winner = 0
	for i := range winCombos {
		// If any of the squares is empty, they don't count
		if b.board[winCombos[i][0]].Byte() == 0 {
			continue
		}
		if b.board[winCombos[i][1]].Byte() == 0 {
			continue
		}
		if b.board[winCombos[i][2]].Byte() == 0 {
			continue
		}

		// Now see if all the values are equal
		if (b.board[winCombos[i][0]].Byte() == b.board[winCombos[i][1]].Byte()) &&
			(b.board[winCombos[i][0]].Byte() == b.board[winCombos[i][2]].Byte()) {
			winner = b.board[winCombos[i][0]].Byte()
			win = true
			Trace.Printf("%v wins\n", winner)
			return
		}
	}
	Trace.Println("no winner, continue")
	return
}

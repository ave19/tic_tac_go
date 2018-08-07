package tictacgo

import (
	"bytes"
	//"strings"
	"fmt"
	"strconv"
)

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
	return b
}

func (b Board) String() string {
	var output bytes.Buffer
	output.WriteString(b.board[0].String())
	output.WriteString("|")
	output.WriteString(b.board[1].String())
	output.WriteString("|")
	output.WriteString(b.board[2].String())
	output.WriteString("\n-+-+-\n")
	output.WriteString(b.board[3].String())
	output.WriteString("|")
	output.WriteString(b.board[4].String())
	output.WriteString("|")
	output.WriteString(b.board[5].String())
	output.WriteString("\n-+-+-\n")
	output.WriteString(b.board[6].String())
	output.WriteString("|")
	output.WriteString(b.board[7].String())
	output.WriteString("|")
	output.WriteString(b.board[8].String())
	output.WriteString("\n")
	return output.String()
}

// Base3 returns the base 3 string for the board state
func (b Board) Base3() string {
	var output bytes.Buffer
	output.WriteString(strconv.FormatInt(int64(b.board[0].Byte()), 10))
	output.WriteString(strconv.FormatInt(int64(b.board[1].Byte()), 10))
	output.WriteString(strconv.FormatInt(int64(b.board[2].Byte()), 10))
	output.WriteString(strconv.FormatInt(int64(b.board[3].Byte()), 10))
	output.WriteString(strconv.FormatInt(int64(b.board[4].Byte()), 10))
	output.WriteString(strconv.FormatInt(int64(b.board[5].Byte()), 10))
	output.WriteString(strconv.FormatInt(int64(b.board[6].Byte()), 10))
	output.WriteString(strconv.FormatInt(int64(b.board[7].Byte()), 10))
	output.WriteString(strconv.FormatInt(int64(b.board[8].Byte()), 10))
	return fmt.Sprintf("%09s", output.String())
}

// Int returns the integer for the board state
func (b Board) Int() uint64 {
	// turn the squares into a string
	x, _ := strconv.ParseUint(b.Base3(), 3, 64)
	return x
}

// SetFromInt you give it an integer, it makes a board out of it
func (b *Board) SetFromInt(newValue uint64) {
	if newValue > 19862 {
		fmt.Println("new value is too big")
	} else {
		b.SetFromBase3(strconv.FormatUint(newValue, 3))
	}
}

// SetFromBase3 you give it a string like '010212120'
func (b *Board) SetFromBase3(newValue string) {
	strRepresentation := fmt.Sprintf("%09s", newValue)

	ul, _ := strconv.ParseInt(strRepresentation[0:1], 10, 8)
	uc, _ := strconv.ParseInt(strRepresentation[1:2], 10, 8)
	ur, _ := strconv.ParseInt(strRepresentation[2:3], 10, 8)
	ml, _ := strconv.ParseInt(strRepresentation[3:4], 10, 8)
	mc, _ := strconv.ParseInt(strRepresentation[4:5], 10, 8)
	mr, _ := strconv.ParseInt(strRepresentation[5:6], 10, 8)
	ll, _ := strconv.ParseInt(strRepresentation[6:7], 10, 8)
	lc, _ := strconv.ParseInt(strRepresentation[7:8], 10, 8)
	lr, _ := strconv.ParseInt(strRepresentation[8:9], 10, 8)

	b.board[0].SetFromByte(byte(ul))
	b.board[1].SetFromByte(byte(uc))
	b.board[2].SetFromByte(byte(ur))
	b.board[3].SetFromByte(byte(ml))
	b.board[4].SetFromByte(byte(mc))
	b.board[5].SetFromByte(byte(mr))
	b.board[6].SetFromByte(byte(ll))
	b.board[7].SetFromByte(byte(lc))
	b.board[8].SetFromByte(byte(lr))
}

// ListEmptySquares still available
func (b Board) ListEmptySquares() (emptySquares []byte) {
	for i := uint(0); i < b.NumSquares; i++ {
		if b.board[i].Byte() == 0 {
			emptySquares = append(emptySquares, byte(i))
		} else {
			//fmt.Println("square ", i, " taken by player ", b.board[i])
		}
	}
	return
}

// Move adjusts the board to a moved state
func (b *Board) Move(square uint8, mark byte) {
	fmt.Println("board.Move(", square, ", ", mark, ")")
	b.board[square].SetFromByte(mark)
}

// Winner return the winning player if there is one
func (b *Board) Winner() (win bool, winner byte) {
	win = false
	winner = 0
	for i := range winCombos {
		//fmt.Println("testing ", winCombos[i])
		if b.board[winCombos[i][0]].Byte() == 0 {
			//fmt.Println(winCombos[i][0], " is 0, continue")
			continue
		}
		if b.board[winCombos[i][1]].Byte() == 0 {
			//fmt.Println(winCombos[i][0], " is 0, continue")
			continue
		}
		if b.board[winCombos[i][2]].Byte() == 0 {
			//fmt.Println(winCombos[i][0], " is 0, continue")
			continue
		}
		if (b.board[winCombos[i][0]].Byte() == b.board[winCombos[i][1]].Byte()) &&
			(b.board[winCombos[i][0]].Byte() == b.board[winCombos[i][2]].Byte()) {
			winner = b.board[winCombos[i][0]].Byte()
			win = true
			//fmt.Println("winner: ", winner)
			return
		}
		//fmt.Println("no winner, continue")
	}
	return
}

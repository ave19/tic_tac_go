package tictacgo

import (
	//"io/ioutil"
	"math"
	"os"
	"testing"
)

func init() {
	InitLogging(os.Stdout, os.Stdout, os.Stdout, os.Stdout)
	//InitLogging(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)
}

func TestStringZero(t *testing.T) {
	testBoard := NewBoard()
	var expected = " | | \n-+-+-\n | | \n-+-+-\n | | \n"
	result := testBoard.String()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestStringNonZero(t *testing.T) {
	testBoard := NewBoard()
	testBoard.board[0] = Square{1}
	testBoard.board[1] = Square{0}
	testBoard.board[2] = Square{2}
	testBoard.board[3] = Square{2}
	testBoard.board[4] = Square{1}
	testBoard.board[5] = Square{2}
	testBoard.board[6] = Square{0}
	testBoard.board[7] = Square{1}
	testBoard.board[8] = Square{0}
	var expected = "X| |O\n-+-+-\nO|X|O\n-+-+-\n |X| \n"
	result := testBoard.String()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestBase3Zero(t *testing.T) {
	testBoard := NewBoard()
	expected := "000000000"
	result := testBoard.Base3()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestBase3NonZero(t *testing.T) {
	testBoard := NewBoard()
	testBoard.board[0] = Square{1}
	testBoard.board[1] = Square{0}
	testBoard.board[2] = Square{2}
	testBoard.board[3] = Square{2}
	testBoard.board[4] = Square{1}
	testBoard.board[5] = Square{2}
	testBoard.board[6] = Square{0}
	testBoard.board[7] = Square{1}
	testBoard.board[8] = Square{0}
	expected := "102212010"
	result := testBoard.Base3()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestIntZero(t *testing.T) {
	testBoard := NewBoard()
	var expected uint64
	result := testBoard.Int()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestIntNonZero(t *testing.T) {
	testBoard := NewBoard()
	testBoard.board[0] = Square{1}
	testBoard.board[1] = Square{0}
	testBoard.board[2] = Square{2}
	testBoard.board[3] = Square{2}
	testBoard.board[4] = Square{1}
	testBoard.board[5] = Square{2}
	testBoard.board[6] = Square{0}
	testBoard.board[7] = Square{1}
	testBoard.board[8] = Square{0}
	var expected uint64 = 8643
	result := testBoard.Int()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromInt(t *testing.T) {
	testBoard := NewBoard()
	var expected uint64 = 8643
	testBoard.SetFromInt(expected)
	result := testBoard.Int()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetTooBigFromInt(t *testing.T) {
	testBoard := NewBoard()
	var expected = uint64(math.Pow(3, 9))
	err := testBoard.SetFromInt(expected)
	result := testBoard.Int()
	if err != nil {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromBase3(t *testing.T) {
	testBoard := NewBoard()
	var base3 = "102212010"
	var expected uint64 = 8643
	testBoard.SetFromBase3(base3)
	result := testBoard.Int()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestMoveX(t *testing.T) {
	testBoard := NewBoard()
	var move byte = 1
	var square uint8 = 4
	testBoard.Move(square, move)
	var expected uint64 = 81
	result := testBoard.Int()
	//fmt.Println(testBoard.String())
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestWinner0(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("000000000")
	//fmt.Println(testBoard.String())
	expected := false
	expectedPlayer := byte(0)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestWinnerTopRowX(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("111000000")
	//fmt.Println(testBoard.String())
	expected := true
	expectedPlayer := byte(1)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestWinnerTopRowO(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("222000000")
	//fmt.Println(testBoard.String())
	expected := true
	expectedPlayer := byte(2)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestLoserTopRow(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("121000000")
	//fmt.Println(testBoard.String())
	expected := false
	expectedPlayer := byte(0)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestWinnerMiddleRowO(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("000222000")
	//fmt.Println(testBoard.String())
	expected := true
	expectedPlayer := byte(2)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestWinnerBottomRowO(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("000000222")
	//fmt.Println(testBoard.String())
	expected := true
	expectedPlayer := byte(2)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestWinnerLeftColumnO(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("200200200")
	//fmt.Println(testBoard.String())
	expected := true
	expectedPlayer := byte(2)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestWinnerMiddleColumnO(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("020020020")
	//fmt.Println(testBoard.String())
	expected := true
	expectedPlayer := byte(2)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestWinnerRightColumnO(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("002002002")
	//fmt.Println(testBoard.String())
	expected := true
	expectedPlayer := byte(2)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestWinnerDownDiagnolO(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("200020002")
	//fmt.Println(testBoard.String())
	expected := true
	expectedPlayer := byte(2)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

func TestWinnerUpDiagnolO(t *testing.T) {
	testBoard := NewBoard()
	testBoard.SetFromBase3("002020200")
	//fmt.Println(testBoard.String())
	expected := true
	expectedPlayer := byte(2)
	result, player := testBoard.Winner()
	if result != expected {
		t.Errorf("Wrong win state: expected '%v', result '%v'", expected, result)
	}
	if player != expectedPlayer {
		t.Errorf("Wront player: expected '%v', result '%v'", expected, result)
	}
}

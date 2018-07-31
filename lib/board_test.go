package tictacgo

import (
	"fmt"
	"testing"
)

func TestStringZero(t *testing.T) {
	testBoard := Board{}
	var expected = " | | \n-+-+-\n | | \n-+-+-\n | | \n"
	result := test_board.String()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestStringNonZero(t *testing.T) {
	testBoard := Board{
		Upper_left:    Square{1},
		Upper_center:  Square{0},
		Upper_right:   Square{2},
		Middle_left:   Square{2},
		Middle_center: Square{1},
		Middle_right:  Square{2},
		Lower_left:    Square{0},
		Lower_center:  Square{1},
		Lower_right:   Square{0},
	}
	var expected = "X| |O\n-+-+-\nO|X|O\n-+-+-\n |X| \n"
	result := test_board.String()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestBase3Zero(t *testing.T) {
	testBoard := Board{}
	expected := "000000000"
	result := test_board.Base3()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestBase3NonZero(t *testing.T) {
	testBoard := Board{
		Upper_left:    Square{1},
		Upper_center:  Square{0},
		Upper_right:   Square{2},
		Middle_left:   Square{2},
		Middle_center: Square{1},
		Middle_right:  Square{2},
		Lower_left:    Square{0},
		Lower_center:  Square{1},
		Lower_right:   Square{0},
	}
	expected := "102212010"
	result := test_board.Base3()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestIntZero(t *testing.T) {
	testBoard := Board{}
	var expected uint64
	result := test_board.Int()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestIntNonZero(t *testing.T) {
	testBoard := Board{
		Upper_left:    Square{1},
		Upper_center:  Square{0},
		Upper_right:   Square{2},
		Middle_left:   Square{2},
		Middle_center: Square{1},
		Middle_right:  Square{2},
		Lower_left:    Square{0},
		Lower_center:  Square{1},
		Lower_right:   Square{0},
	}
	var expected uint64 = 8643
	result := test_board.Int()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromInt(t *testing.T) {
	testBoard := Board{}
	var expected uint64 = 8643
	test_board.SetFromInt(expected)
	result := test_board.Int()
	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromBase3(t *testing.T) {
	testBoard := Board{}
	var base3 = "102212010"
	var expected uint64 = 8643
	test_board.SetFromBase3(base3)
	result := test_board.Int()

	if result != expected {
		t.Errorf("Test failed! expected '%v', result '%v'", expected, result)
	}
}

func TestMoveX(t *testing.T) {
	testBoard := Board{}
	var move byte = 1
	var square = 1
	test_board.Addresses()
	test_board.Move(square, move)
	//s := test_board.GetSquareByInt(1)
	b := test_board.Upper_center.Byte()
	fmt.Printf("b: %v\n", b)
	fmt.Printf("board\n%v\n", test_board.String())
	test_board.Addresses()
}

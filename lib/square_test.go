package tic_tac_go

import (
	"testing"
)

func TestSetFromStringX(t *testing.T) {
	test_square := Square{}
	expected := "X"
	test_square.SetFromString(expected)
	result := test_square.String()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromStringO(t *testing.T) {
	test_square := Square{}
	expected := "O"
	test_square.SetFromString(expected)
	result := test_square.String()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromStringEmpty(t *testing.T) {
	test_square := Square{}
	expected := " "
	test_square.SetFromString(expected)
	result := test_square.String()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

func TestByteX(t *testing.T) {
	test_square := Square{}
	var expected byte = 1
	var test_string string = "X"
	test_square.SetFromString(test_string)
	result := test_square.Byte()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromByte(t *testing.T) {
	test_square := Square{}
	var expected byte = 1
	var test_byte byte = 1
	test_square.SetFromByte(test_byte)
	result := test_square.Byte()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

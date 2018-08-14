package tictacgo

import (
	//"os"
	"io/ioutil"
	"testing"
)

func init() {
	//InitLogging(os.Stdout, os.Stdout, os.Stdout, os.Stdout)
	InitLogging(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)
}

func TestSetFromStringX(t *testing.T) {
	testSquare := Square{}
	expected := "X"
	testSquare.SetFromString(expected)
	result := testSquare.String()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromStringO(t *testing.T) {
	testSquare := Square{}
	expected := "O"
	testSquare.SetFromString(expected)
	result := testSquare.String()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromStringEmpty(t *testing.T) {
	testSquare := Square{}
	expected := " "
	testSquare.SetFromString(expected)
	result := testSquare.String()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

func TestByteX(t *testing.T) {
	testSquare := Square{}
	var expected byte = 1
	var testString = "X"
	testSquare.SetFromString(testString)
	result := testSquare.Byte()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromByte(t *testing.T) {
	testSquare := Square{}
	var expected byte = 1
	var testByte byte = 1
	testSquare.SetFromByte(testByte)
	result := testSquare.Byte()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

func TestSetFromByteBad(t *testing.T) {
	testSquare := Square{}
	var testByte byte = 3
	err := testSquare.SetFromByte(testByte)
	if err != nil {
		t.Errorf("Expected failure!")
	}
}

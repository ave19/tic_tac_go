package tictacgo

import (
	"testing"
)

func TestSetName(t *testing.T) {
	testPlayer := NewPlayer()
	expected := "Bob"
	testPlayer.SetName(expected)
	result := testPlayer.Name()

	if result != expected {
		t.Errorf("Setting the square failed! expected '%v', result '%v'", expected, result)
	}
}

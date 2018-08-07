package tictacgo

import (
	"fmt"
	"math/rand"
)

//RememberLosingStrategy avoids things it picked when it lost
type RememberLosingStrategy struct {
	enabled  bool
	name     string
	memories map[uint][]byte
	moves    map[uint][]byte
}

//NewRememberLosingStrategy constructor
func NewRememberLosingStrategy() *RememberLosingStrategy {
	return &RememberLosingStrategy{
		name:     "RememberLosingStrategy",
		enabled:  true,
		memories: make(map[uint][]byte),
		moves:    make(map[uint][]byte),
	}
}

// Name of the strategy
func (r *RememberLosingStrategy) Name() string {
	return r.name
}

// Move avoids things it picked when it lost
func (r *RememberLosingStrategy) Move(b Board) byte {
	var index int
	var choice byte
	emptySquares := b.ListEmptySquares()
	//fmt.Println("emptySquares = ", emptySquares)
	//fmt.Println("emptySquareLen = ", len(emptySquares))

	// compare to the memories
	boardNumber := uint(b.Int())
	//fmt.Println("board number: ", boardNumber)
	//fmt.Println("memories about this board: ", r.memories[boardNumber])
	availableSquares := filterSquares(emptySquares, r.memories[boardNumber])

	//fmt.Println("available squares:", availableSquares)

	if len(availableSquares) != 0 {
		index = rand.Intn(len(availableSquares))
		choice = byte(availableSquares[index])
	} else {
		//fmt.Println("* There are no available squares, falling back to random")
		// reset the memory?
		r.memories[boardNumber] = nil
		if len(emptySquares) != 0 {
			index = rand.Intn(len(emptySquares))
			choice = byte(emptySquares[index])
		} else {
			choice = 0
		}
	}
	fmt.Println("choosing square ", choice)

	// Now we jot down our move so we can think about it later
	r.moves[boardNumber] = append(r.moves[boardNumber], choice)
	return choice
}

func isBanned(potential byte, banned []byte) bool {
	for _, b := range banned {
		if potential == b {
			return true
		}
	}
	return false
}

func filterSquares(potential []byte, banned []byte) (accepted []byte) {
	for _, p := range potential {
		if isBanned(p, banned) {
			// don't add to the list
		} else {
			accepted = append(accepted, p)
		}
	}
	return accepted
}

func uniqueSquares(list []byte) (newList []byte) {
	set := make(map[byte]struct{})
	for _, value := range list {
		set[value] = struct{}{}
	}
	for key := range set {
		newList = append(newList, key)
	}
	return
}

// Win adjust strategy on a win
func (r *RememberLosingStrategy) Win(b Board) {
	winningBoard := int(b.Int())
	fmt.Println("winningBoard = ", winningBoard)
	fmt.Println(r.moves)
	return
}

// Lose adjust strategy on a lose
func (r *RememberLosingStrategy) Lose(b Board) {
	losingBoard := int(b.Int())
	fmt.Println("losingBoard = ", losingBoard)
	fmt.Println(r.moves)
	for board, moves := range r.moves {
		fmt.Println(" -- board = ", board)
		fmt.Println(" ---- moves = ", moves)
		for _, m := range r.moves[board] {
			r.memories[board] = append(r.memories[board], m)
			fmt.Println(" ----- m:", m)
			fmt.Println(" >> new memory, board:", board, "moves:", r.memories[board])
		}
		r.memories[board] = uniqueSquares(r.memories[board])
		delete(r.moves, board)
	}
	fmt.Println("new memories: ", r.memories)
	return
}

// Draw adjust strategy on a draw
func (r *RememberLosingStrategy) Draw(b Board) {
	fmt.Println(r.moves)
	return
}

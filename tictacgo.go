package main

import (
	"fmt"
	"github.com/ave19/tictacgo/lib"
	//"io/ioutil"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	fmt.Println("WOULD YOU LIKE TO PLAY A GAME")

	moveChan := make(chan byte)

	playerOne := tictacgo.NewPlayer()
	playerOne.SetName("Alice")
	playerOne.SetNumber(1)
	playerOne.SetStrategy(tictacgo.NewRememberWinningStrategy())
	bChanOne := make(chan tictacgo.Board)
	go playerOne.Play(bChanOne, moveChan)

	playerTwo := tictacgo.NewPlayer()
	playerTwo.SetName("Bob")
	playerTwo.SetNumber(2)
	playerTwo.SetStrategy(tictacgo.NewRememberLosingStrategy())
	bChanTwo := make(chan tictacgo.Board)
	go playerTwo.Play(bChanTwo, moveChan)

	var bChannels []chan tictacgo.Board
	bChannels = append(bChannels, bChanOne)
	bChannels = append(bChannels, bChanTwo)

	players := []*tictacgo.Player{&playerOne, &playerTwo}
	maxGames := 10000

	for game := 0; game < maxGames; game++ {
		b := tictacgo.NewBoard()
		for _, p := range players {
			p.Stats()
		}
		fmt.Println("\nGame ", game+1)

		roundOver := false
		for count := 0; !roundOver; count++ {
			fmt.Println("------------------------------------------")
			playerSelector := count % len(players)
			player := players[playerSelector]
			fmt.Println(" - Current Player: ", player.Name())
			fmt.Println("Round ", count, ":")
			fmt.Println(b.Int())
			fmt.Println(b.String())
			bChannels[playerSelector] <- b
			b.Move(<-moveChan, player.Number())
			if count >= int(b.NumSquares) {
				roundOver = true
			}
			win, mark := b.Winner()
			if win {
				fmt.Println("==============================================")
				for _, p := range players {
					if p.Number() == mark {
						fmt.Println("Winner: ", p.Name())
						p.Win(b)
					} else {
						fmt.Println("Loser: ", p.Name())
						p.Lose(b)
					}
				}
				roundOver = true
			} else {
				remainingSquares := len(b.ListEmptySquares())
				if remainingSquares == 0 {
					fmt.Println("It was a tie.")
					roundOver = true
					for _, p := range players {
						p.Draw(b)
					}
				}
			}
		}
		fmt.Println("Final Board:")
		fmt.Println(b.Int())
		fmt.Println(b.String())
	}

	for _, p := range players {
		p.Stats()
	}
}

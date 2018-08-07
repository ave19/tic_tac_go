package main

import (
	"fmt"
	"github.com/ave19/tictacgo/lib"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("WOULD YOU LIKE TO PLAY A GAME")
	playerOne := tictacgo.NewPlayer()
	playerOne.SetName("Alice")
	playerOne.SetNumber(1)

	playerTwo := tictacgo.NewPlayer()
	playerTwo.SetName("Bob")
	playerTwo.SetNumber(2)
	playerTwo.SetStrategy(tictacgo.NewRememberLosingStrategy())

	fmt.Println("Players:")

	players := []*tictacgo.Player{&playerOne, &playerTwo}
	for _, p := range players {
		fmt.Println("Player ", p.Number(), ":  ", p.Name())
		fmt.Println("  Strategy: ", p.Strategy().Name())
	}

	maxGames := 10000
	//roundOver := false
	//var round []byte

	for game := 0; game < maxGames; game++ {
		b := tictacgo.NewBoard()
		for _, p := range players {
			p.Stats()
		}
		fmt.Println("\nGame ", game+1)
		roundOver := false
		for count := 0; !roundOver; count++ {
			fmt.Println("Round ", count, ":")
			fmt.Println(b.Int())
			fmt.Println(b.String())
			playerSelector := count % len(players)
			player := players[playerSelector]
			fmt.Println(" - Current Player: ", player.Name())
			b.Move(player.Move(b), player.Number())
			if count >= int(b.NumSquares) {
				roundOver = true
			}
			win, mark := b.Winner()
			if win {
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

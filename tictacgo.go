package main

import (
	"fmt"
	"github.com/ave19/tictacgo/lib"
)

func main() {
	fmt.Println("WOULD YOU LIKE TO PLAY A GAME")

	/*
		fmt.Printf("as string -> '%v'\n", a.String())
		fmt.Printf("as state -> '%v'\n", a.Byte())

		a.SetFromString("X")
		fmt.Printf("as string -> '%v'\n", a.String())
		fmt.Printf("as state -> '%v'\n", a.Byte())

		fmt.Println("board:")
		fmt.Println(b.String())
		fmt.Printf("board upper left: %v\n", b.Upper_left.Byte())

		b.Middle_left.SetFromString("X")
		b.Lower_center.SetFromString("O")
		fmt.Println("board2:")
		fmt.Println(b.String())

		fmt.Printf("board base 3: %v\n", b.Base3())
		fmt.Printf("board:\n%v", b.String())
		fmt.Printf("board int: %v\n", b.Int())
		fmt.Println("-----------------------------------------")
		fmt.Printf("board reset:\n")
		b.SetFromInt(0)
		fmt.Printf("base 3: %v\n", b.Base3())
		fmt.Printf("board:\n%v", b.String())
		fmt.Printf("int: %v\n", b.Int())
		fmt.Printf("middle: %v\n", b.GetSquareByInt(4).String())

		fmt.Println("-----------------------------------------")
		fmt.Println("setting board to 249")
		b.SetFromInt(249)
		fmt.Printf("base 3: %v\n", b.Base3())
		fmt.Printf("board:\n%v", b.String())
		fmt.Printf("board int: %v\n", b.Int())
		fmt.Printf("middle: %v\n", b.GetSquareByInt(4).String())

		fmt.Println("-----------------------------------------")
		fmt.Println("setting board to 12654")
		b.SetFromInt(12654)
		fmt.Printf("base 3: %v\n", b.Base3())
		fmt.Printf("board:\n%v", b.String())
		fmt.Printf("board int: %v\n", b.Int())
		fmt.Printf("middle: %v\n", b.GetSquareByInt(4).String())

		fmt.Println("-----------------------------------------")
		fmt.Println("setting board to 8643")
		b.SetFromInt(8643)
		fmt.Printf("base 3: %v\n", b.Base3())
		fmt.Printf("board:\n%v", b.String())
		fmt.Printf("board int: %v\n", b.Int())
		fmt.Printf("middle: %v\n", b.GetSquareByInt(4).String())
		b.Move(1, '1')
		fmt.Printf("moved board:\n%v", b.String())


		fmt.Println("-----------------------------------------")
		fmt.Println("setting board to 6561")
		b.SetFromInt(6561)
		fmt.Printf("base 3: %v\n", b.Base3())
		fmt.Printf("board:\n%v", b.String())
		fmt.Printf("board int: %v\n", b.Int())
		fmt.Printf("middle: %v\n", b.GetSquareByInt(4).String())

		fmt.Println("-----------------------------------------")
		b.SetFromInt(1236561)
		fmt.Println("-----------------------------------------")
	*/

	b := tictacgo.NewBoard()
	b.Move(3, 2)
	fmt.Println(b.String())
	win, winner := b.Winner()
	fmt.Println("win: ", win)
	fmt.Println("winner: ", winner)

}

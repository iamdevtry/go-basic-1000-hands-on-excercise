package main

import (
	"bufio"
	"fmt"
	"os"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {
	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	fmt.Printf("Devtry's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
> list 	: lists all the games
> quit	: quits
`)
		if !in.Scan() {
			break
		}
		fmt.Println()
		switch in.Text() {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		}

	}
}

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var nplayers int
var depthAI int
var turnAI bool
var showAnalysis bool
var sleepAI time.Duration
var clearScreen bool

func init() {
	rand.Seed(time.Now().Unix())
	flag.IntVar(&nplayers, "n", 1, "number of players")
	flag.IntVar(&depthAI, "d", 5, "AI depth (-1 for best play)")
	flag.BoolVar(&turnAI, "c", true, "computer starts")
	flag.BoolVar(&showAnalysis, "a", true, "show computer analysis")
	flag.BoolVar(&clearScreen, "l", true, "show one board at a time (clear screen)")
	flag.DurationVar(&sleepAI, "s", 5e8, `simulate thinking by "sleeping"`)
}

func scanInt() (in int) {
	fmt.Printf("> ")
	fmt.Scanf("%d", &in)
	return in
}

func main() {
	flag.Parse()
	game := NewGame()
	if !turnAI {
		fmt.Print(game.board)
	}
	for !game.Over() {
		switch {
		case nplayers >= 2 || nplayers == 1 && !turnAI:
			if ok := game.Move(scanInt()); !ok {
				continue
			}
		default:
			time.Sleep(sleepAI)
			game.MoveAI(depthAI)
		}

		fmt.Println(game.board)
		if showAnalysis {
			w, d, l := game.analyze(-1)
			fmt.Printf(" %s%v%s", green, w, clrRst)
			fmt.Printf(" %v", d)
			fmt.Printf(" %s%v%s\n\n", red, l, clrRst)
		}

		turnAI = !turnAI
	}

	if w := game.Winner(); w != None {
		fmt.Printf("\nPlayer %s won\n", w)
	} else {
		fmt.Printf("\nDraw\n")
	}
	time.Sleep(sleepAI) // useful if program run in a loop, WarGames style
}

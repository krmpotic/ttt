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
var sleepAI time.Duration

func init() {
	rand.Seed(time.Now().Unix())
	flag.IntVar(&nplayers, "n", 1, "number of players")
	flag.IntVar(&depthAI, "d", 5, "AI depth (-1 for best play)")
	flag.BoolVar(&turnAI, "c", true, "computer starts")
	flag.BoolVar(&ShowAnalysis, "a", true, "show computer analysis")
	flag.BoolVar(&ClearScreen, "l", true, "show one board at a time (clear screen)")
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

	fmt.Println(game)
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

		fmt.Println(game)

		turnAI = !turnAI
	}

	if w := game.Winner(); w != None {
		fmt.Printf("Player %s won\n", w)
	} else {
		fmt.Printf("Draw\n")
	}
	time.Sleep(sleepAI) // useful if program run in a loop, WarGames style
}

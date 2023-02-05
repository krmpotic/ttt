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
	flag.BoolVar(&turnAI, "c", false, "computer starts")
	flag.BoolVar(&showAnalysis, "a", false, "show computer analysis")
	flag.BoolVar(&clearScreen, "cls", true, "show one board at a time (clear screen)")
	flag.DurationVar(&sleepAI, "s", 0, `simulate thinking by "sleeping"`)
}

type game struct {
	turn  player
	board board
}

func newGame() *game {
	return &game{turn: x}
}

func (g *game) Move(n int) (ok bool) {
	if n < 0 || n > 8 || g.board[n] != none || g.Over() {
		return false
	}

	g.board[n], g.turn = g.turn, g.turn.other()
	return true
}

func (g *game) unMove(n int) {
	g.board[n], g.turn = none, g.turn.other()
}

func (g *game) MoveAI(depth int) (ok bool) {
	w, r, l := g.analyze(depth)
	var m int
	switch {
	case len(w) > 0:
		m = w[rand.Intn(len(w))]
	case len(r) > 0:
		m = r[rand.Intn(len(r))]
	case len(l) > 0:
		m = l[rand.Intn(len(l))]
	default:
		return false
	}
	g.Move(m)
	return true
}

func (g *game) analyze(depth int) (wins, rest, losses []int) {
	if g.Over() {
		return nil, nil, nil
	}

	moves := g.board.moves()
	if depth == 0 {
		return nil, moves, nil
	}
	for _, m := range moves {
		g.Move(m)
		switch {
		case g.board.won():
			wins = append(wins, m)
		case g.board.full():
			rest = append(rest, m)
		default:
			w, r, l := g.analyze(depth - 1) // enemy
			switch {
			case len(w) > 0:
				losses = append(losses, m)
			case len(r) > 0:
				rest = append(rest, m)
			case len(l) > 0:
				wins = append(wins, m)
			}
		}
		g.unMove(m)
	}
	return
}

func (g *game) Over() (gameOver bool) {
	return g.board.full() || g.board.won()
}

func (g *game) Winner() player {
	if !g.board.won() {
		return none
	}
	return g.turn.other()
}

func scanInt() (in int) {
	fmt.Printf("> ")
	fmt.Scanf("%d", &in)
	return in
}

func main() {
	flag.Parse()
	game := newGame()
	if !turnAI {
		fmt.Print(game)
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
		fmt.Print(game)
		turnAI = !turnAI
	}

	if w := game.Winner(); w != none {
		fmt.Printf("\nPlayer %s won\n", w)
	} else {
		fmt.Printf("\nDraw\n")
	}
}

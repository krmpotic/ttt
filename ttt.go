package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var nplayers int
var turnAI bool

type player int
type board [9]player
type game struct {
	turn  player
	board board
}

const (
	O    = player(-1)
	none = player(0)
	X    = player(1)
)

const (
	Red    = "\033[1;31m"
	Green  = "\033[1;32m"
	ClrRst = "\033[0m"
)

func init() {
	rand.Seed(time.Now().Unix())
	flag.IntVar(&nplayers, "n", 1, "number of players")
	flag.BoolVar(&turnAI, "c", false, "computer starts")
}

func NewGame() game {
	return game{turn: X}
}

func (p player) other() player {
	return -1 * p
}

func (p player) String() string {
	switch p {
	case X:
		return Green + "X" + ClrRst
	case O:
		return Red + "O" + ClrRst
	default:
		return " "
	}
}

func (b board) full() bool {
	for _, p := range b {
		if p == none {
			return false
		}
	}
	return true
}

func (b board) won() bool {
	return false ||
		// rows
		b[0] != none && b[0] == b[1] && b[0] == b[2] ||
		b[3] != none && b[3] == b[4] && b[3] == b[5] ||
		b[6] != none && b[6] == b[7] && b[6] == b[8] ||
		// columns
		b[0] != none && b[0] == b[3] && b[0] == b[6] ||
		b[1] != none && b[1] == b[4] && b[1] == b[7] ||
		b[2] != none && b[2] == b[5] && b[2] == b[8] ||
		// diagonals
		b[0] != none && b[0] == b[4] && b[0] == b[8] ||
		b[2] != none && b[2] == b[4] && b[2] == b[6]
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

func (g *game) MoveAI() (ok bool) {
	w, d, l := g.Analyze()
	switch {
	case len(w) > 0:
		g.Move(w[rand.Intn(len(w))])
	case len(d) > 0:
		g.Move(d[rand.Intn(len(d))])
	case len(l) > 0:
		g.Move(l[rand.Intn(len(l))])
	default:
		return false
	}
	return true
}

func (g *game) Analyze() (wins []int, draws []int, losses []int) {
	if g.Over() {
		return nil, nil, nil
	}
	for m := range g.board {
		if ok := g.Move(m); !ok {
			continue
		}
		switch {
		case g.board.won():
			wins = append(wins, m)
		case g.board.full():
			draws = append(draws, m)
		default:
			w, d, l := g.Analyze() // enemy
			switch {
			case len(w) > 0:
				losses = append(losses, m)
			case len(d) > 0:
				draws = append(draws, m)
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

func (g *game) Board() (s string) {
	// show field number if empty
	f := func(i int) string {
		if g.board[i] == none {
			return fmt.Sprintf("%d", i)
		}
		return g.board[i].String()
	}
	s += "\n"
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(0), f(1), f(2))
	s += fmt.Sprintf("═══╬═══╬═══\n")
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(3), f(4), f(5))
	s += fmt.Sprintf("═══╬═══╬═══\n")
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(6), f(7), f(8))
	s += "\n"
	return s
}

func main() {
	flag.Parse()
	game := NewGame()
	if !turnAI {
		fmt.Println(game.Board())
	}
	for !game.Over() {
		switch {
		case nplayers >= 2 || nplayers == 1 && !turnAI:
			var move int
			fmt.Printf("> ")
			fmt.Scanf("%d", &move)
			if ok := game.Move(move); !ok {
				continue
			}
		default:
			time.Sleep(time.Second)
			game.MoveAI()
		}
		fmt.Print(game.Board())
		fmt.Println(game.Analyze())
		turnAI = !turnAI
	}

	if w := game.Winner(); w != none {
		fmt.Printf("\nPlayer %s won\n", w)
	} else {
		fmt.Printf("\nDraw\n")
	}
}

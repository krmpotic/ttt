package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var nplayers int
var turnAI bool
var showAnalysis bool
var random bool
var sleepAI time.Duration

const (
	Red    = "\033[1;31m"
	Green  = "\033[1;32m"
	ClrRst = "\033[0m"
)

func init() {
	rand.Seed(time.Now().Unix())
	flag.IntVar(&nplayers, "n", 1, "number of players")
	flag.BoolVar(&turnAI, "c", false, "computer starts")
	flag.BoolVar(&random, "r", false, "AI plays random moves")
	flag.BoolVar(&showAnalysis, "a", false, "show computer analysis")
	flag.DurationVar(&sleepAI, "s", 0, `simulate thinking by "sleeping"`)
}

type game struct {
	turn  player
	board board
}

func NewGame() *game {
	return &game{turn: X}
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
	if random {
		return g.moveRand()
	}
	return g.moveBest()
}

func (g *game) moveBest() (ok bool) {
	w, d, l := g.Analyze()
	var m int
	switch {
	case len(w) > 0:
		m = w[rand.Intn(len(w))]
	case len(d) > 0:
		m = d[rand.Intn(len(d))]
	case len(l) > 0:
		m = l[rand.Intn(len(l))]
	default:
		return false
	}
	g.Move(m)
	return true
}

func (g *game) moveRand() (ok bool) {
	m_ := g.board.moves()
	if len(m_) == 0 {
		return false
	}

	g.Move(m_[rand.Intn(len(m_))])
	return true
}

func (g *game) Analyze() (wins []int, draws []int, losses []int) {
	if g.Over() {
		return nil, nil, nil
	}
	m_ := g.board.moves()
	for m := range m_ {
		g.Move(m)
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

func (g *game) String() (s string) {
	s = fmt.Sprintf("\n%s", g.board)
	if showAnalysis {
		w, d, l := g.Analyze()
		s += fmt.Sprintf(" %s%v%s", Green, w, ClrRst)
		s += fmt.Sprintf(" %v", d)
		s += fmt.Sprintf(" %s%v%s", Red, l, ClrRst)
		s += "\n"
	}
	return s
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
			game.MoveAI()
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

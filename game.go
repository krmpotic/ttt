package main

import (
	"fmt"
	"math/rand"
)

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

func (g *game) String() (s string) {
	s = fmt.Sprintf("\n%s", g.board)
	if showAnalysis {
		w, d, l := g.analyze(-1)
		s += fmt.Sprintf(" %s%v%s", green, w, clrRst)
		s += fmt.Sprintf(" %v", d)
		s += fmt.Sprintf(" %s%v%s", red, l, clrRst)
		s += "\n"
	}
	return s
}

package main

import (
	"math/rand"
)

type game struct {
	turn  Player
	board board
}

func NewGame() *game {
	return &game{turn: X}
}

func (g *game) Move(n int) (ok bool) {
	if n < 0 || n > 8 || g.board[n] != None || g.Over() {
		return false
	}

	g.board[n], g.turn = g.turn, g.turn.Other()
	return true
}

func (g *game) unMove(n int) {
	g.board[n], g.turn = None, g.turn.Other()
}

func (g *game) MoveAI(depth int) (ok bool) {
	w, r, l := g.analyze(depth)
	rand := func(m []int) int {
		return m[rand.Intn(len(m))]
	}
	var m int
	switch {
	case len(w) > 0:
		m = rand(w)
	case len(r) > 0:
		m = rand(r)
	case len(l) > 0:
		m = rand(l)
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

func (g *game) Winner() Player {
	if !g.board.won() {
		return None
	}
	return g.turn.Other()
}

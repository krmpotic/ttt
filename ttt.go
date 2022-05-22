package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type board [9]player
type game struct {
	board board
	turn  player
}

type player int

const (
	O    = player(-1)
	none = player(0)
	X    = player(1)
)

func NewGame() game {
	return game{turn: X}
}

func (p player) Other() player {
	if p == X {
		return O
	}
	return X
}

func (g *game) Move(n int) (ok bool) {
	if n < 0 || n > 8 ||
		g.board[n] != 0 || g.Over() {
		return false
	}

	g.board[n] = g.turn
	g.turn = g.turn.Other()
	return true
}

func (g *game) unMove(n int) {
	g.board[n] = none
	g.turn = g.turn.Other()
}

func (g *game) MoveAI() {
	w, d, l := g.analyze()
	var move int
	switch {
	case len(w) > 0:
		move = w[rand.Intn(len(w))]
	case len(d) > 0:
		move = d[rand.Intn(len(d))]
	case len(l) > 0:
		move = l[rand.Intn(len(l))]
	}
	g.Move(move)
}

// analyze returns which moves for the player who's turn it is
// win & draw & lose
func (g *game) analyze() (wins []int, draws []int, losses []int) {
	if g.Over() {
		return nil, nil, nil
	}

	for m := 0; m < 9; m++ {
		if ok := g.Move(m); !ok {
			continue
		}
		if g.Won() {
			wins = append(wins, m)
			g.unMove(m)
			continue
		}
		if g.Over() {
			draws = append(draws, m)
			g.unMove(m)
			continue
		}
		w, d, l := g.analyze() // enemy
		if len(w) > 0 {
			losses = append(losses, m)
		} else if len(d) > 0 {
			draws = append(draws, m)
		} else if len(l) > 0 {
			wins = append(wins, m)
		}
		g.unMove(m)
	}
	return
}

func (g *game) Over() (gameOver bool) {
	return g.board.Full() || g.Won()
}

func (g *game) Won() bool {
	b := g.board
	return false ||
		// rows
		b[0] != 0 && b[0] == b[1] && b[0] == b[2] ||
		b[3] != 0 && b[3] == b[4] && b[3] == b[5] ||
		b[6] != 0 && b[6] == b[7] && b[6] == b[8] ||
		// columns
		b[0] != 0 && b[0] == b[3] && b[0] == b[6] ||
		b[1] != 0 && b[1] == b[4] && b[1] == b[7] ||
		b[2] != 0 && b[2] == b[5] && b[2] == b[8] ||
		// diagonals
		b[0] != 0 && b[0] == b[4] && b[0] == b[8] ||
		b[2] != 0 && b[2] == b[4] && b[2] == b[6]
}

func (g *game) Winner() player {
	if !g.Won() {
		return none
	}
	// rely on keeping track of turns
	return g.turn.Other()
}

func (b board) Full() bool {
	for i := 0; i < 9; i++ {
		if b[i] == 0 {
			return false
		}
	}
	return true
}
func (b board) String() (s string) {
	// show field number if empty
	f := func(i int) string {
		if b[i] == 0 {
			return fmt.Sprintf("%d", i)
		}
		return b[i].String()
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

func (p player) String() string {
	const (
		Red   = "\033[1;31m"
		Green = "\033[1;32m"
		Reset = "\033[0m"
	)

	switch p {
	case X:
		return Green + "X" + Reset
	case O:
		return Red + "O" + Reset
	default:
		return " "
	}

}

func main() {
	game := NewGame()

	fmt.Print(game.board)
	for !game.Over() {
		fmt.Println(game.analyze())
		var move int
		fmt.Printf("> ")
		fmt.Scanf("%d", &move)

		if move == 9 { // easter-egg
			game.MoveAI()
			goto easterskip
		}

		if ok := game.Move(move); !ok {
			continue
		}
	easterskip:
		game.MoveAI()

		fmt.Print(game.board)

	}

	if w := game.Winner(); w != none {
		fmt.Printf("\nPlayer %s won\n", w)
	}
}

package main

import (
	"fmt"
)

const (
	empty = 0
	X     = 1
	O     = 2
)

type field int
type board [3][3]field
type turn int

func (b board) Done() bool {
	return false ||
		// rows
		b[0][0] != empty && b[0][0] == b[0][1] && b[0][0] == b[0][2] ||
		b[1][0] != empty && b[1][0] == b[1][1] && b[1][0] == b[1][2] ||
		b[2][0] != empty && b[2][0] == b[2][1] && b[2][0] == b[2][2] ||
		// columns
		b[0][0] != empty && b[0][0] == b[1][0] && b[0][0] == b[2][0] ||
		b[0][1] != empty && b[0][1] == b[1][1] && b[0][1] == b[2][1] ||
		b[0][2] != empty && b[0][2] == b[1][2] && b[0][2] == b[2][2] ||
		// diagonals
		b[0][0] != empty && b[0][0] == b[1][1] && b[0][0] == b[2][2] ||
		b[0][2] != empty && b[0][2] == b[1][1] && b[0][2] == b[2][0]
}

func (f field) String() string {
	switch f {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}

func (b board) String() (s string) {
	s += fmt.Sprintf(" %s | %s | %s\n", b[0][0], b[0][1], b[0][2])
	s += fmt.Sprintf("---+---+---\n")
	s += fmt.Sprintf(" %s | %s | %s \n", b[1][0], b[1][1], b[1][2])
	s += fmt.Sprintf("---+---+---\n")
	s += fmt.Sprintf(" %s | %s | %s \n", b[2][0], b[2][1], b[2][2])
	return s
}

func main() {
	var b board

	turn := turn(X)
	fmt.Println(b)
	for !b.Done() {
		var row, col int
		fmt.Scanf("%d %d", &row, &col)
		b[row][col] = field(turn)
		turn.next()
		fmt.Println(b)
	}
}

func (t *turn) next() {
	if *t == turn(X) {
		*t = turn(O)
	} else {
		*t = turn(X)
	}
}

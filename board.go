package main

import "fmt"

type board [9]player

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

func (b board) moves() (m []int) {
	for i, p := range b {
		if p == none {
			m = append(m, i)
		}
	}
	return m
}

func (b board) String() (s string) {
	// show field number if empty
	f := func(i int) string {
		if b[i] == none {
			return fmt.Sprintf("%d", i)
		}
		return b[i].String()
	}
	if clearScreen {
		s += clrScr
	}
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(0), f(1), f(2))
	s += fmt.Sprintf("═══╬═══╬═══\n")
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(3), f(4), f(5))
	s += fmt.Sprintf("═══╬═══╬═══\n")
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(6), f(7), f(8))
	return s
}

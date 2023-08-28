package main

import (
	"fmt"
	"slices"
)

type board [9]player

func (b board) full() bool {
	return !slices.Contains(b[:], none)
}

func (b board) won() bool {
	s := func(i, j, k int) bool {
		return b[i] != none && b[i] == b[j] && b[i] == b[k]
	}

	return s(0, 4, 8) || s(2, 4, 6) || // diagonals
		s(0, 1, 2) || s(3, 4, 5) || s(6, 7, 8) || // rows
		s(0, 3, 6) || s(1, 4, 7) || s(2, 5, 8) // colums
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

package main

import (
	"fmt"
	"slices"
)

var ClearScreen bool

type Board [9]Player

func (b Board) Full() bool {
	return !slices.Contains(b[:], None)
}

func (b Board) Won() bool {
	eq := func(i, j, k int) bool {
		return b[i] != None && b[i] == b[j] && b[i] == b[k]
	}

	return eq(0, 4, 8) || eq(2, 4, 6) || // diagonals
		eq(0, 1, 2) || eq(3, 4, 5) || eq(6, 7, 8) || // rows
		eq(0, 3, 6) || eq(1, 4, 7) || eq(2, 5, 8) // colums
}

func (b Board) Moves() (m []int) {
	for i, p := range b {
		if p == None {
			m = append(m, i)
		}
	}
	return m
}

func (b Board) String() (s string) {
	// show field number if empty
	f := func(i int) string {
		if b[i] == None {
			return fmt.Sprintf("%d", i)
		}
		return b[i].String()
	}
	if ClearScreen {
		s += clrScr
	}
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(0), f(1), f(2))
	s += fmt.Sprintf("═══╬═══╬═══\n")
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(3), f(4), f(5))
	s += fmt.Sprintf("═══╬═══╬═══\n")
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(6), f(7), f(8))
	return s
}

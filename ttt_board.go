package main

const (
	o    = player(-1)
	none = player(0)
	x    = player(1)
)

type player int

func (p player) other() player {
	return -1 * p
}

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

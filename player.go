package main

type player int

const (
	red    = "\033[1;31m"
	green  = "\033[1;32m"
	clrRst = "\033[0m"
	clrScr = "\033[H\033[2J"
)

const (
	o    = player(-1)
	none = player(0)
	x    = player(1)
)

func (p player) other() player {
	return -1 * p
}

func (p player) String() string {
	switch p {
	case x:
		return green + "X" + clrRst
	case o:
		return red + "O" + clrRst
	default:
		return " "
	}
}

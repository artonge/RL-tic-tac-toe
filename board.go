package main

import "fmt"

type caseType int
type boardLine []caseType
type board []boardLine

const (
	e caseType = iota
	x
	o
)

// Create a new empty board
func newBoard() board {
	return board{
		boardLine{e, e, e},
		boardLine{e, e, e},
		boardLine{e, e, e},
	}
}

// Return true if two board are equals
func (b board) isEqual(b2 board) bool {
	for i, l := range b {
		for j := range l {
			if b[i][j] != b2[i][j] {
				return false
			}
		}
	}
	return true
}

// Return true if the board is full
func (b board) isFull() bool {
	for _, l := range b {
		for _, c := range l {
			if c == e {
				return false
			}
		}
	}
	return true
}

func (b board) getWinnerSign() caseType {
	// Check all lines and columns
	for i, l := range b {
		// Lines
		if l[0] == l[1] && l[1] == l[2] {
			return l[0]
		}
		// Columns
		if b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return b[0][i]
		}
	}
	// Check diagonals
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return b[0][0]
	}
	if b[2][0] == b[1][1] && b[1][1] == b[0][2] {
		return b[2][0]
	}
	return e
}

func (b board) copy() board {
	newB := newBoard()
	copy(newB[0], b[0])
	copy(newB[1], b[1])
	copy(newB[2], b[2])
	return newB
}

func (b board) String() string {
	return fmt.Sprintf("%v\n%v\n%v", b[0], b[1], b[2])
}

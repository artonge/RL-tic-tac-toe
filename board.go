package main

import "fmt"

// ---
// BOARDCASE DEFINITION
// ---
type boardCase int

const (
	e boardCase = iota
	x
	o
)

func (c boardCase) String() string {
	switch c {
	case e:
		return "_"
	case x:
		return "x"
	case o:
		return "o"
	}
	return ""
}

// ---
// BOARDLINE DEFINITION
// ---
type boardLine []boardCase

// ---
// BOARD DEFINITION
// ---
type board []boardLine

// Create a new empty board
func newBoard() board {
	return board{
		boardLine{e, e, e},
		boardLine{e, e, e},
		boardLine{e, e, e},
	}
}

// Count the number of diff between two boards
func (b board) diff(b2 board) int {
	diff := 0
	for i, l := range b {
		for j := range l {
			if b[i][j] != b2[i][j] {
				diff++
			}
		}
	}
	return diff
}

// Return a copy of the board rotated by 90°
func (b board) rotate() board {
	bRot := newBoard()
	// Transpose the board
	for i, row := range b {
		for j := range row {
			bRot[i][j] = b[j][i]
		}
	}
	// Reverse each rows
	for _, row := range bRot {
		tmp := row[0]
		row[0] = row[2]
		row[2] = tmp
	}
	return bRot
}

// Return the number of occurence a sign in the board
func (b board) countSign(s boardCase) int {
	count := 0
	for _, l := range b {
		for _, c := range l {
			if c == s {
				count++
			}
		}
	}
	return count
}

// Return true if the board is full
func (b board) isFull() bool {
	return b.countSign(e) == 0
}

// Return the winner
// Need three signs aligned
// If no winner, return e
func (b board) getWinnerSign() boardCase {
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

// Copy a board to a new board
func (b board) copy() board {
	newB := newBoard()
	copy(newB[0], b[0])
	copy(newB[1], b[1])
	copy(newB[2], b[2])
	return newB
}

// Overide default String method
func (b board) String() string {
	return fmt.Sprintf("%v\n%v\n%v", b[0], b[1], b[2])
}

// Stringify a board
// Example:
// X - X
// - O -
// - - -
// ==> "X-X--O---"
func (b board) serialize() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v", b[0][0], b[0][1], b[0][2], b[1][0], b[1][1], b[1][2], b[2][0], b[2][1], b[2][2])
}

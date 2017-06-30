package main

import (
	"fmt"
)

func main() {
	// Create the board and two agent
	var b board
	a1 := agent{history: make(map[string]*state), gameMoves: make(map[string]*state), sign: x}
	a2 := agent{history: make(map[string]*state), gameMoves: make(map[string]*state), sign: o}
	a1Win := 0
	a2Win := 0
	// Set the number of games to play and when a1 forget and stop to learn
	loopNb := 10000
	for i := 0; i < loopNb; i++ {
		b = newBoard()
		if i >= loopNb/2 {
			// After half, always make a1 forget
			a1.history = make(map[string]*state)
			// At half, reset the agent's win count
			if i == loopNb/2 {
				// Display stats
				fmt.Println("-------------")
				fmt.Println("Both learning")
				fmt.Println("-------------")
				fmt.Printf("%v wins %v%% times\n", a1.sign, float64(a1Win)/float64(i)*100)
				fmt.Printf("%v wins %v%% times\n", a2.sign, float64(a2Win)/float64(i)*100)
				fmt.Println("====================================")
				a1Win = 0
				a2Win = 0
			}
		}
		// Play the game until the board is full or there is a winner
		// Alternate who begin to play
		for true {
			if i%2 == 0 {
				if b.isFull() || b.getWinnerSign() != e {
					break
				}
				a1.play(b)
				if b.isFull() || b.getWinnerSign() != e {
					break
				}
				a2.play(b)
			} else {
				if b.isFull() || b.getWinnerSign() != e {
					break
				}
				a2.play(b)
				if b.isFull() || b.getWinnerSign() != e {
					break
				}
				a1.play(b)
			}
		}
		// Get the winner and give rewards
		winnerSign := b.getWinnerSign()
		if winnerSign == a1.sign {
			a1.feed(1)
			a2.feed(0)
			a1Win++
		} else if winnerSign == a2.sign {
			a1.feed(0)
			a2.feed(1)
			a2Win++
		} else {
			a1.feed(0)
			a2.feed(0)
		}
	}
	// Display new stats
	fmt.Println("------------------")
	fmt.Println("x forget, o learns")
	fmt.Println("------------------")
	fmt.Printf("%v wins %v%% times\n", a1.sign, float64(a1Win)/float64(loopNb/2)*100)
	fmt.Printf("%v wins %v%% times\n", a2.sign, float64(a2Win)/float64(loopNb/2)*100)
}

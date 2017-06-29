package main

import "fmt"

func main() {
	var b board
	a1 := agent{sign: x}
	a2 := agent{sign: o}
	loopNb := 100000
	for i := 0; i < loopNb; i++ {
		b = newBoard()
		// At half, reset agent 1 and reset agent 2 wins count
		if i == loopNb/2 {
			fmt.Printf("%v win %v times\n", a1.sign, a1.wins)
			fmt.Printf("%v win %v times\n", a2.sign, a2.wins)
			a1 = agent{sign: x}
			a2.wins = 0
		}
		for true {
			a1.play(b)
			if b.isFull() {
				break
			}
			a2.play(b)
		}
		winnerSign := b.getWinnerSign()
		// fmt.Printf("%v	Winner is %v\n", b, winnerSign)
		if winnerSign == a1.sign {
			a1.win()
			a2.lose()
		} else if winnerSign == a2.sign {
			a1.lose()
			a2.win()
		} else {
			a1.lose()
			a2.lose()
		}
	}
	fmt.Printf("%v win %v times\n", a1.sign, a1.wins)
	fmt.Printf("%v win %v times\n", a2.sign, a2.wins)
}

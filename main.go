package main

import (
	"fmt"
)

func main() {
	var b board
	a1 := agent{history: map[string]*state{}, gameMoves: map[string]*state{}, sign: x}
	a2 := agent{history: map[string]*state{}, gameMoves: map[string]*state{}, sign: o}
	loopNb := 1
	for i := 0; i < loopNb; i++ {
		b = newBoard()
		// At half, reset agent 1 and reset agent 2 wins count
		if i == loopNb*(2/3) && false {
			fmt.Printf("%v win %v times\n", a1.sign, a1.wins)
			fmt.Printf("%v win %v times\n", a2.sign, a2.wins)
			a1 = agent{history: map[string]*state{}, gameMoves: map[string]*state{}, sign: x}
			a2.wins = 0
		}
		//fmt.Println(i, len(a2.history))
		for true {
			if b.isFull() || b.getWinnerSign() != e {
				break
			}
			a1.play(b)
			if b.isFull() || b.getWinnerSign() != e{
				break
			}
			a2.play(b)
		}
		winnerSign := b.getWinnerSign()
		//fmt.Printf("%v	Winner is %v\n---\n", b, winnerSign)
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
	fmt.Println(len(a1.history))
	for _, s := range a1.history {
		if s.boardState.countSign(e) == 8 {
			fmt.Println(*s)
		}
	}
	fmt.Println(len(a2.history))
	for _, s := range a2.history {
		if s.boardState.countSign(e) == 8 {
			fmt.Println(*s)
		}
	}

}

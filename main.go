package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

func nbWins(wins []entry, c boardCase) int {
	count := 0
	for _, e := range wins {
		if e.Value == c {
			count++
		}
	}
	return count
}

type entry struct {
	Timestamp int64     `json:"timestamp"`
	Value     boardCase `json:"value"`
}

func (e entry) String() string {
	return fmt.Sprintf("%v, %v\n", e.Timestamp, e.Value)
}

func main() {
	start := time.Now().UnixNano()
	plot := len(os.Args) > 1 && os.Args[1] == "--plot"
	// Create the board and two agent
	var b board
	a1 := agent{history: make(map[string]*state), gameMoves: make(map[string]*state), sign: x}
	a2 := agent{history: make(map[string]*state), gameMoves: make(map[string]*state), sign: o}
	// Set the number of games to play and when a1 forget and stop to learn
	loopNb := 3000
	wins := make([]entry, 0, loopNb)
	interWins := make([]entry, 0, loopNb/2)
	for i := 0; i < loopNb; i++ {
		b = newBoard()
		if i >= loopNb/2 {
			// After half, always make a1 forget
			// At half, reset the agent's win count
			if i == loopNb/2 {
				a2.history = make(map[string]*state)
				// Display stats
				fmt.Println("-------------")
				fmt.Println("Both learning")
				fmt.Println("-------------")
				fmt.Printf("%v wins %v%% times\n", a1.sign, float64(nbWins(interWins, a1.sign))/float64(i)*100)
				fmt.Printf("%v wins %v%% times\n", a2.sign, float64(nbWins(interWins, a2.sign))/float64(i)*100)
				fmt.Println("====================================")
				interWins = make([]entry, 0, loopNb/2)
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
			wins = append(wins, entry{time.Now().UnixNano() - start, winnerSign})
			interWins = append(interWins, entry{time.Now().UnixNano() - start, winnerSign})
		} else if winnerSign == a2.sign {
			a1.feed(0)
			a2.feed(1)
			wins = append(wins, entry{time.Now().UnixNano() - start, winnerSign})
			interWins = append(interWins, entry{time.Now().UnixNano() - start, winnerSign})
		} else {
			a1.feed(0)
			a2.feed(0)
		}
	}
	// Display new stats
	fmt.Println("------------------")
	fmt.Println("x forget, o learns")
	fmt.Println("------------------")
	fmt.Printf("%v wins %v%% times\n", a1.sign, float64(nbWins(interWins, a1.sign))/float64(loopNb/2)*100)
	fmt.Printf("%v wins %v%% times\n", a2.sign, float64(nbWins(interWins, a2.sign))/float64(loopNb/2)*100)

	if plot {
		generateFigure(wins, loopNb, a1, a2)
	}
}

// Generate figure from the wins array
func generateFigure(wins []entry, loopNb int, a1 agent, a2 agent) {
	// Create plot
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	// Set plot meta data
	p.Title.Text = "Both learn then O forget"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Number of wins"
	// Build plot data
	ptsX := make(plotter.XYs, nbWins(wins, a1.sign)+1)
	ptsO := make(plotter.XYs, nbWins(wins, a2.sign)+1)
	countX := 0
	countO := 0
	for _, w := range wins[:loopNb] {
		if w.Value == x {
			countX++
			ptsX[countX].Y = float64(countX)
			ptsX[countX].X = float64(w.Timestamp)
		} else if w.Value == o {
			countO++
			ptsO[countO].Y = float64(countO)
			ptsO[countO].X = float64(w.Timestamp)
		}
	}
	// Add data to plot
	err = plotutil.AddLines(p, "X", ptsX, "O", ptsO)
	if err != nil {
		panic(err)
	}
	// Save the plot to a PNG file.
	err = p.Save(4*vg.Inch, 4*vg.Inch, "points.png")
	if err != nil {
		panic(err)
	}
}

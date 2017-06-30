package main

import (
	"math/rand"
	"time"
	"fmt"
)

// Associate to a board state a value
type state struct {
	boardState board
	value      float64
	count      int
}

func (s state) String() string {
	return fmt.Sprintf("%v	{value: %v - count: %v}\n---", s.boardState, s.value, s.count)
}

// Describe an action of the agent
// aka where the agent put an x
type action struct {
	i, j int
}

// An agent that will learn to play tic-tac-toe
type agent struct {
	history    map[string]*state // States list of all the games
	gameMoves  map[string]*state // States list of the current game
	sign       caseType // The sign use to play
	wins       int      // Number of wins
}

var r *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
}

// The policy
func (a *agent) policy(b board) action {
	// fmt.Println("	Running policy")
	maxVal := -1000.0
	var nextAction action
	posStates, posActions := a.possibleNextStates(b)
	// Get the hight value next state
	for i, s := range posStates {
		if s.value > maxVal {
			maxVal = s.value
			nextAction = posActions[i]
		}
	}
	// 10% of the time explore other actions
	if r.Intn(100) < 10 {
		nextAction = posActions[r.Intn(len(posActions))]
	}
	return nextAction
}

// Return the possible states from a given states
func (a *agent) possibleNextStates(b board) ([]*state, []action) {
	// fmt.Println("	Getting possible next states from:")
	posStates := []*state{}
	posActions := []action{}
	var posBoard board
	// Look for empty cases on the board
	for i, l := range b {
		for j, c := range l {
			if c == e {
				// When a case is empty, add the associated state to posStates
				posBoard = b.copy()
				posBoard[i][j] = a.sign
				posStates = append(posStates, a.getState(posBoard))
				posActions = append(posActions, action{i, j})
			}
		}
	}
	return posStates, posActions
}

// Called when we need the agent to play
func (a *agent) play(b board) {
	// fmt.Printf("%v plays\n", a.sign)
	action := a.policy(b)
	s := a.getState(b)
	a.gameMoves[s.boardState.serialize()] = s
	b[action.i][action.j] = a.sign
}

// Called when the agent wins
func (a *agent) win() {
	a.wins++
	a.updateHistory(3)
}

// Called when the agent loses
func (a *agent) lose() {
	a.updateHistory(-1)
}

func (a *agent) updateHistory(reward int) {
	// fmt.Printf("	Update %v's gameMoves\n", a.sign)
	for _, s := range a.gameMoves {
		s.value, s.count = incrementalMean(s.value, reward, s.count)
		fmt.Println(*s)
	}
	a.gameMoves = map[string]*state{}
}

// Return a new mean value given old mean, new a value and the count of values involved
func incrementalMean(oldMean float64, value int, k int) (float64, int) {
	// fmt.Println("	Computing incremental mean")
	return (oldMean*float64(k) + float64(value)) / float64(k+1), k + 1
}

// Return a state from the history or create a new one
func (a *agent) getState(b board) *state {
	 //fmt.Println("	Getting state for board:")
	 //fmt.Println(b)
	var s *state
	for i := 0; i < 4; i++ {
		s, ok := a.history[b.serialize()];
		if ok {
			//fmt.Println("Found state:")
			//fmt.Println(s)
			return s
		}
		b = b.rotate()
	}
	if s == nil {		
		// If state does not existe, create a new one and add it to states
		s = &state{
			value:      0,
			boardState: b.copy(),
		}
		a.history[s.boardState.serialize()] = s		
		//fmt.Println("	None found, creating one")
	}
	return s
}
